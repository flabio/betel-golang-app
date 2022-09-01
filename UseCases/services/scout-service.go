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

	var userDTO dto.ScoutRequest

	userDTO.ChurchId = ChurchId

	if len(userDTO.Email) == 0 {
		userDTO.Email = " "
	}
	if validateBirthDayScout(userDTO, scoutService, context) {
		return
	}
	userToCreated, msg := getMappingScout(userDTO, context)
	if msg != "" {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}

	filename, _ := UploadFile(context)
	documentfile, _ := UploadFileDocument(context)
	userToCreated.Image = filename
	userToCreated.Identification = documentfile
	createdUser, errs := scoutService.IUser.SetInsertUser(userToCreated)
	if errs != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(errs.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildCreatedResponse(createdUser))

}

// update scout
func (scoutService *scoutService) Update(ChurchId uint, context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	var userDTO dto.ScoutRequest
	userDTO.ChurchId = ChurchId

	if err != nil {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if len(userDTO.Email) == 0 {
		userDTO.Email = " "
	}
	if validateBirthDayScout(userDTO, scoutService, context) {
		return

	}
	userToCreated, msg := getMappingScout(userDTO, context)
	if msg != "" {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}

	findById, _ := scoutService.IUser.GetProfileUser(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
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
	u, err := scoutService.IUser.SetEditUser(userToCreated, uint(id))

	if err != nil {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(u))

}

//method private

// validarUser

func validateBirthDayScout(u dto.ScoutRequest, scoutService *scoutService, context *gin.Context) bool {

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

func getMappingScout(userDTO dto.ScoutRequest, context *gin.Context) (entity.User, string) {
	user := entity.User{}
	err := context.ShouldBind(&userDTO)
	if err != nil {
		return user, utilities.GetMsgErrorRequired(err)
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		return user, err.Error()
	}
	return user, ""
}
