package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"

	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

var wg sync.WaitGroup

type userService struct {
	IUser Interfaces.IUser
}

//NewUserService creates a new instance of UserService
func NewUserService() InterfacesService.IUserService {
	return &userService{
		IUser: repositorys.NewUserRepository(),
	}
}

//List user
func (userService *userService) GetListUserService(context *gin.Context) {

	users, err := userService.IUser.GetAllUser()

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "OK", users))
}

// list ListKingsScouts

func (userService *userService) GetListKingsScoutsService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	users, err := userService.IUser.GetListKingsScouts(uint(id))

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "OK", users))
}

//Create user
func (userService *userService) SetCreateService(context *gin.Context) {
	// userChan := make(chan int)
	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.UserDTO

	context.ShouldBind(&userDTO)

	if validarUser(userDTO, userService, context, constantvariables.OPTION_CREATE) {
		return
	}

	smapping.FillStruct(&userToCreated, smapping.MapFields(&userDTO))

	userToCreated.Password = utilities.HashAndSalt([]byte(userToCreated.Password))

	filename, _ := UploadFile(context)

	userToCreated.Image = filename
	createdUser, errs := userService.IUser.SetInsertUser(userToCreated)
	if errs != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errs.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, createdUser)
	context.JSON(http.StatusOK, res)

	roleToCreated.RolId = userDTO.RolId
	roleToCreated.UserId = uint(createdUser.Id)

	err := userService.IUser.SetInsertRole(roleToCreated)
	if err != nil {
		log.Println(err)
		_, err := userService.IUser.SetRemoveUser(uint(createdUser.Id))
		if err != nil {
			res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NOT_DELETED)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

}

//update user
func (userService *userService) SetUpdateService(context *gin.Context) {
	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.UserDTO

	context.ShouldBind(&userDTO)

	if validarUser(userDTO, userService, context, constantvariables.OPTION_EDIT) {
		return
	}

	smapping.FillStruct(&userToCreated, smapping.MapFields(&userDTO))

	roleToCreated.RolId = userDTO.RolId
	roleToCreated.UserId = userDTO.Id

	wg.Add(1)
	go goRunitaUpdateRole(userService, roleToCreated)
	wg.Wait()

	findById, _ := userService.IUser.GetProfileUser(uint(userDTO.Id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
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
	u, err := userService.IUser.SetEditUser(userToCreated)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, u))

}
func (userService *userService) SetUpdatePasswordService(context *gin.Context) {
	user := entity.User{}
	var userDTO dto.UserPasswordDTO

	errDTO := context.ShouldBind(&userDTO)
	if errDTO != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errDTO.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	checkError(err)
	errp := userService.IUser.SetChangePassword(user)
	if errp != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_PASSWORD_UPDATE, nil)
	context.JSON(http.StatusOK, res)
}
func (userService *userService) SetRemoveService(context *gin.Context) {
	chanels := make(chan bool)
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	user, errprofile := userService.IUser.GetProfileUser(uint(id))
	if errprofile != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	go func() {
		err := userService.IUser.SetRemoveRoleUser(user.Id)
		if err != nil {
			chanels <- false
			return
		}
		chanels <- true
		close(chanels)
	}()
	if <-chanels {

		_, err := userService.IUser.SetRemoveUser(user.Id)
		if err != nil {
			res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NOT_DELETED)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, user)
		context.JSON(http.StatusOK, res)
		return
	}
	res := utilities.BuildNotFoudResponse()
	context.JSON(http.StatusBadRequest, res)
}
func (userService *userService) GetProfileService(Id uint, context *gin.Context) {

	user, err := userService.IUser.GetProfileUser(Id)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", user)
	context.JSON(http.StatusOK, res)

}

func (userService *userService) GetFindUserService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	user, err := userService.IUser.GetProfileUser(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", user)
	context.JSON(http.StatusOK, res)

}
func (userService *userService) GetFindUserNameLastNameService(context *gin.Context) {

	search := context.Param("search")

	users, err := userService.IUser.GetFindUserNameLastName(search)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "OK", users))
	return

}
func (userService *userService) GetAllService(context *gin.Context) {
	total := userService.IUser.GetCountUser()
	var limit int64 = 9
	page, begin := utilities.Pagination(context, int(limit))
	pages := (total / limit)
	if (total % limit) != 0 {
		pages++
	}

	users, err := userService.IUser.GetPaginationUsers(begin, int(limit))

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	context.JSON(http.StatusOK, struct {
		Data  []entity.User `json:"data"`
		Total int64         `json:"total"`
		Page  int           `json:"page"`
		Pages int64         `json:"pages"`
		Limit int64         `json:"limit"`
	}{
		Data:  users,
		Total: total,
		Page:  page,
		Pages: pages,
		Limit: limit,
	})
}

//method private
//goRunitaCreateRole

func goRunitaCreateRole(userService *userService, roleToCreated entity.Role) {
	wg.Done()
	err := userService.IUser.SetInsertRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
	}
}

//goRunitaUpdateRole
func goRunitaUpdateRole(userService *userService, roleToCreated entity.Role) {
	wg.Done()
	role, err := userService.IUser.SetEditRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
	}

	if role.Id == 0 {
		userService.IUser.SetInsertRole(roleToCreated)
	}
}

//validarUser
func validarUser(u dto.UserDTO, userService *userService, context *gin.Context, option int) bool {
	context.ShouldBind(&u)
	switch option {
	case 1:
		existEmail, _ := userService.IUser.IsDuplicateEmail(u.Email)
		existsIdentification := userService.IUser.IsDuplicateIdentificatio(u.Identification)
		if existsIdentification {
			context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.IDENTIFICATION_EXIST))
			return true
		}
		if existEmail {
			context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.EMAIL_EXIST))

			return true
		}
		if validarUserField(u, context) {
			return true
		}

	case 2:
		if u.Id == 0 {
			context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID))
			return true
		}
		if validarUserField(u, context) {
			return true
		}
	}

	return false
}

func validarUserField(u dto.UserDTO, context *gin.Context) bool {

	if len(u.Name) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NAME))

		return true
	}
	if len(u.LastName) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.LAST_NAME))

		return true
	}
	if len(u.Identification) == 0 {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.IDENTIFICATION))

		return true
	}

	if len(u.TypeIdentification) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.TYPE_IDENTIFICATION))

		return true
	}
	if len(u.Birthplace) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.BIRTH_PLACE))

		return true
	}
	if len(u.Gender) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GENDER))

		return true
	}
	if len(u.CivilStatus) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.CIVIL_STATUS))

		return true
	}
	if len(u.Email) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.EMAIL))

		return true
	}

	if u.RolId == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ROL))

		return true
	}

	if u.ChurchId == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.CHURCH_ID))
		return true
	}

	if u.Password == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PASSWORD))
		return true
	}
	if u.Password != u.ConfirmPassword {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PASSWORD_CONFIRM))

		return true
	}
	return false
}
