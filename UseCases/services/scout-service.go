package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"
	"strconv"
	"strings"
	"time"

	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

var wgs sync.WaitGroup

type scoutService struct {
	IUser Interfaces.IUser
}

// NewUserService creates a new instance of UserService
func NewScoutService() InterfacesService.IScoutService {
	return &scoutService{
		IUser: repositorys.NewUserRepository(),
	}
}

// list ListKingsScouts

func (scoutService *scoutService) ListKingsScouts(id uint, context *gin.Context) {

	users, err := scoutService.IUser.GetListKingsScouts(uint(id))

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(users))
}

// Create scout
func (scoutService *scoutService) Create(ChurchId uint, context *gin.Context) {
	userChan := make(chan int)
	option := 0

	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.ScoutDTO

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
		userToCreated.Identification = documentfile
		createdUser, errs := scoutService.IUser.SetInsertUser(userToCreated)
		if errs != nil {
			res := utilities.BuildErrResponse(errs.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		res := utilities.BuildCreatedResponse(createdUser)
		context.JSON(http.StatusOK, res)
		userChan <- int(createdUser.Id)
		close(userChan)
	}(userChan)

	select {
	case user_id := <-userChan:

		roleToCreated.RolId = 29
		roleToCreated.UserId = uint(user_id)

		err := scoutService.IUser.SetInsertRole(roleToCreated)
		if err != nil {
			log.Println(err)
			_, err := scoutService.IUser.SetRemoveUser(uint(user_id))
			if err != nil {
				res := utilities.BuildErrResponse(constantvariables.NOT_DELETED)
				context.JSON(http.StatusBadRequest, res)
				return
			}
			res := utilities.BuildErrResponse(err.Error())
			context.JSON(http.StatusBadRequest, res)
			return
		}
	}
}

// update scout
func (scoutService *scoutService) Update(ChurchId uint, context *gin.Context) {
	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.ScoutDTO

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

	findById, _ := scoutService.IUser.GetProfileUser(uint(userDTO.Id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
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

	if len(findById.Identification) == 0 {
		userToCreated.Identification = documentfile
	} else {
		if documentfile != "" {
			userToCreated.Identification = documentfile
		} else {
			userToCreated.Image = findById.Image
			userToCreated.Identification = findById.Identification
		}
	}
	u, err := scoutService.IUser.SetEditUser(userToCreated)

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildUpdatedResponse(u)
	context.JSON(http.StatusOK, res)

}

//method private
//goRunitaCreateRole

func goRunitaCreateScoutRole(scoutService *scoutService, roleToCreated entity.Role) {
	wgs.Done()
	err := scoutService.IUser.SetInsertRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
		return
	}
}

// goRunitaUpdateRole
func goRunitaUpdateScoutRole(scoutService *scoutService, roleToCreated entity.Role) {
	wgs.Done()
	role, err := scoutService.IUser.SetEditRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
		return
	}

	if role.Id == 0 {
		scoutService.IUser.SetInsertRole(roleToCreated)
	}
}

// validarUser
func validarScout(u dto.ScoutDTO, scoutService *scoutService, context *gin.Context, option int) bool {
	context.ShouldBind(&u)

	if len(u.Name) == 0 {
		res := utilities.BuildErrResponse(constantvariables.NAME)
		context.JSON(http.StatusBadRequest, res)
		return true
	}
	if len(u.LastName) == 0 {
		res := utilities.BuildErrResponse(constantvariables.LAST_NAME)
		context.JSON(http.StatusBadRequest, res)
		return true
	}
	if len(u.Identification) == 0 {

		res := utilities.BuildErrResponse(constantvariables.IDENTIFICATION)
		context.JSON(http.StatusBadRequest, res)
		return true
	}

	if len(u.TypeIdentification) == 0 {
		res := utilities.BuildErrResponse(constantvariables.TYPE_IDENTIFICATION)
		context.JSON(http.StatusBadRequest, res)
		return true
	}
	if len(u.Birthplace) == 0 {
		res := utilities.BuildErrResponse(constantvariables.BIRTH_PLACE)
		context.JSON(http.StatusBadRequest, res)
		return true
	}
	if len(u.Gender) == 0 {
		res := utilities.BuildErrResponse(constantvariables.GENDER)
		context.JSON(http.StatusBadRequest, res)
		return true
	}
	if u.ChurchId == 0 {
		res := utilities.BuildErrResponse(constantvariables.CHURCH_ID)
		context.JSON(http.StatusBadRequest, res)
		return true
	}
	if u.SubDetachmentId == 0 {
		res := utilities.BuildErrResponse(constantvariables.DETACHMENT_ID)
		context.JSON(http.StatusBadRequest, res)
		return true
	}
	if option == 1 {

		if u.Id == 0 {
			res := utilities.BuildErrResponse(constantvariables.ID)
			context.JSON(http.StatusBadRequest, res)
			return true
		}
	} else {

		existsIdentification := scoutService.IUser.IsDuplicateIdentificatio(u.Identification)
		if existsIdentification {

			res := utilities.BuildErrResponse(constantvariables.IDENTIFICATION_EXIST)
			context.JSON(http.StatusBadRequest, res)
			return true
		}

	}
	return false
}
func validateBirthDayScout(u dto.ScoutDTO, scoutService *scoutService, context *gin.Context) bool {

	context.ShouldBind(&u)
	YearBirthday := strings.Split(u.Birthday, "-")
	currentDate := time.Now()

	yearB, _ := strconv.ParseUint(YearBirthday[0], 0, 0)

	difYear := uint64(currentDate.Year()) - yearB
	res := utilities.BuildErrResponse(constantvariables.BIRTH_PLACE_CHECK)

	if u.SubDetachmentId == 1 {
		if difYear < 3 || difYear > 8 {
			context.JSON(http.StatusBadRequest, res)
			return true
		}
		return false
	}
	if u.SubDetachmentId == 2 {
		if difYear <= 8 || difYear > 11 {
			context.JSON(http.StatusBadRequest, res)
			return true
		}
		return false
	}
	if u.SubDetachmentId == 3 {
		if difYear <= 11 || difYear > 14 {
			context.JSON(http.StatusBadRequest, res)
			return true
		}
		return false
	}
	if u.SubDetachmentId == 4 {
		if difYear <= 14 || difYear > 18 {
			context.JSON(http.StatusBadRequest, res)
			return true
		}
		return false
	}
	return false
}
