package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"

	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

var wgs sync.WaitGroup

//UserService is a contract.....
type ScoutService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	ListKingsScouts(context *gin.Context)
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

func (scoutService *scoutService) ListKingsScouts(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	users, err := scoutService.userRepository.ListKingsScouts(uint(id))

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "OK", users))
}

//Create scout
func (scoutService *scoutService) Create(context *gin.Context) {
	userChan := make(chan int)
	option := 0

	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.ScoutDTO

	context.ShouldBind(&userDTO)

	if len(userDTO.Email) == 0 {
		userDTO.Email = " "
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
func (scoutService *scoutService) Update(context *gin.Context) {
	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.ScoutDTO

	err := context.ShouldBind(&userDTO)
	if len(userDTO.Email) == 0 {
		userDTO.Email = " "
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
