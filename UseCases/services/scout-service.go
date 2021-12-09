package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"strconv"
	"strings"
	"time"

	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

var wgs sync.WaitGroup

//UserService is a contract.....
type ScoutService interface {
	Create(SubdetachmentId uint, ChurchId uint, context *gin.Context)
	Update(SubdetachmentId uint, ChurchId uint, context *gin.Context)
	ListKingsScouts(id uint, context *gin.Context)
}

type scoutService struct {
	userRepository repositorys.UserRepository
}

//NewUserService creates a new instance of UserService
func NewScoutService() ScoutService {
	var userRepo repositorys.UserRepository = repositorys.NewUserRepository()
	return &scoutService{
		userRepository: userRepo,
	}
}

// list ListKingsScouts

func (scoutService *scoutService) ListKingsScouts(id uint, context *gin.Context) {

	users, err := scoutService.userRepository.ListKingsScouts(uint(id))

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "OK", users))
}

//Create scout
func (scoutService *scoutService) Create(SubdetachmentId uint, ChurchId uint, context *gin.Context) {
	userChan := make(chan int)
	option := 0

	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.ScoutDTO
	userDTO.SubDetachmentId = SubdetachmentId
	userDTO.ChurchId = ChurchId

	context.ShouldBind(&userDTO)

	if len(userDTO.Email) == 0 {
		userDTO.Email = " "
	}
	if validateBirthDayScout(userDTO, scoutService, context) {
		return
	}
	if validarScout(userDTO, scoutService, context, option) {
		return
	}

	smapping.FillStruct(&userToCreated, smapping.MapFields(&userDTO))

	go func(userChan chan<- int) {

		filename, _ := UploadFile(context)
		documentfile, _ := UploadFileDocument(context)
		userToCreated.Image = filename
		userToCreated.DocumentIdentification = documentfile
		createdUser, errs := scoutService.userRepository.InsertUser(userToCreated)
		if errs != nil {
			validadErrors(errs, context)
			return
		}
		res := utilities.BuildCreateResponse(createdUser)
		context.JSON(http.StatusOK, res)
		userChan <- int(createdUser.Id)
		close(userChan)
	}(userChan)

	select {
	case user_id := <-userChan:

		roleToCreated.RolId = 29
		roleToCreated.UserId = uint(user_id)

		err := scoutService.userRepository.InsertRole(roleToCreated)
		if err != nil {
			log.Println(err)
			result, err := scoutService.userRepository.DeleteUser(uint(user_id))
			if err != nil {
				validadErrorRemove(result, context)
				return
			}
			validadErrors(err, context)
			return
		}
	}
}

//update scout
func (scoutService *scoutService) Update(SubdetachmentId uint, ChurchId uint, context *gin.Context) {
	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.ScoutDTO
	userDTO.SubDetachmentId = SubdetachmentId
	userDTO.ChurchId = ChurchId
	err := context.ShouldBind(&userDTO)
	if len(userDTO.Email) == 0 {
		userDTO.Email = " "
	}
	if validateBirthDayScout(userDTO, scoutService, context) {
		return
	}
	if validarScout(userDTO, scoutService, context, 1) {
		return
	}

	smapping.FillStruct(&userToCreated, smapping.MapFields(&userDTO))

	roleToCreated.RolId = 29
	roleToCreated.UserId = userDTO.Id

	wgs.Add(1)
	go goRunitaUpdateScoutRole(scoutService, roleToCreated)
	wgs.Wait()

	findById, _ := scoutService.userRepository.ProfileUser(uint(userDTO.Id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	filename, err := UploadFile(context)

	if len(findById.Image) == 0 {
		userToCreated.Image = filename
	} else {
		if filename != "" {
			userToCreated.Image = filename
		} else {
			userToCreated.Image = findById.Image
		}
	}

	documentfile, _ := UploadFileDocument(context)

	if len(findById.DocumentIdentification) == 0 {
		userToCreated.DocumentIdentification = documentfile
	} else {
		if documentfile != "" {
			userToCreated.DocumentIdentification = documentfile
		} else {
			userToCreated.Image = findById.Image
			userToCreated.DocumentIdentification = findById.DocumentIdentification
		}
	}
	u, err := scoutService.userRepository.EditUser(userToCreated)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(u)
	context.JSON(http.StatusOK, res)

}

//method private
//goRunitaCreateRole

func goRunitaCreateScoutRole(scoutService *scoutService, roleToCreated entity.Role) {
	wgs.Done()
	err := scoutService.userRepository.InsertRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
	}
}

//goRunitaUpdateRole
func goRunitaUpdateScoutRole(scoutService *scoutService, roleToCreated entity.Role) {
	wgs.Done()
	role, err := scoutService.userRepository.EditRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
	}

	if role.Id == 0 {
		scoutService.userRepository.InsertRole(roleToCreated)
	}
}

//validarUser
func validarScout(u dto.ScoutDTO, scoutService *scoutService, context *gin.Context, option int) bool {
	context.ShouldBind(&u)
	msg := utilities.MessageRequired{}

	if len(u.Name) == 0 {
		validadRequiredMsg(msg.RequiredName(), context)
		return true
	}
	if len(u.LastName) == 0 {
		validadRequiredMsg(msg.RequiredLastName(), context)
		return true
	}
	if len(u.Identification) == 0 {

		validadRequiredMsg(msg.RequiredIdentification(), context)
		return true
	}

	if len(u.TypeIdentification) == 0 {
		validadRequiredMsg(msg.RequiredTypeIdentification(), context)
		return true
	}
	if len(u.Birthplace) == 0 {
		validadRequiredMsg(msg.RequiredBirthplace(), context)
		return true
	}
	if len(u.Gender) == 0 {
		validadRequiredMsg(msg.RequiredGender(), context)
		return true
	}
	if u.ChurchId == 0 {
		validadRequiredMsg(msg.RequiredChurch(), context)
		return true
	}
	if u.SubDetachmentId == 0 {
		validadRequiredMsg(msg.RequiredDetachment(), context)
		return true
	}
	if option == 1 {

		if u.Id == 0 {
			validadRequiredMsg(msg.RequiredId(), context)
			return true
		}
	} else {

		existsIdentification := scoutService.userRepository.IsDuplicateIdentificatio(u.Identification)
		if existsIdentification {
			validadRequiredMsg(msg.RequiredExistIdentification(), context)
			return true
		}

	}
	return false
}
func validateBirthDayScout(u dto.ScoutDTO, scoutService *scoutService, context *gin.Context) bool {

	context.ShouldBind(&u)
	msg := utilities.MessageRequired{}
	YearBirthday := strings.Split(u.Birthday, "-")
	currentDate := time.Now()

	yearB, _ := strconv.ParseUint(YearBirthday[0], 0, 0)

	difYear := uint64(currentDate.Year()) - yearB

	if u.SubDetachmentId == 1 {
		if difYear < 3 || difYear > 8 {
			validadBirdateRequiredMsg(msg.RequiredBirthday(), context)
			return true
		}
		return false
	}
	if u.SubDetachmentId == 2 {
		if difYear <= 8 || difYear > 11 {
			validadBirdateRequiredMsg(msg.RequiredBirthday(), context)
			return true
		}
		return false
	}
	if u.SubDetachmentId == 3 {
		if difYear <= 11 || difYear > 14 {
			validadBirdateRequiredMsg(msg.RequiredBirthday(), context)
			return true
		}
		return false
	}
	if u.SubDetachmentId == 4 {
		if difYear <= 14 || difYear > 18 {
			validadBirdateRequiredMsg(msg.RequiredBirthday(), context)
			return true
		}
		return false
	}
	return false
}
