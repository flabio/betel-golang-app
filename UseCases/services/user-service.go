package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"

	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
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

// NewUserService creates a new instance of UserService
func NewUserService() InterfacesService.IUserService {
	return &userService{
		IUser: repositorys.NewUserRepository(),
	}
}

// List user
func (userService *userService) GetListUserService(context *gin.Context) {
	var usersLists []dto.UserResponse
	users, err := userService.IUser.GetAllUser()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, data := range users {
		user := getListsUserDto(data)
		usersLists = append(usersLists, user)
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(usersLists))
}

// list ListKingsScouts

func (userService *userService) GetListKingsScoutsService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	users, err := userService.IUser.GetListKingsScouts(uint(id))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(users))
}

// Create user
func (userService *userService) SetCreateService(context *gin.Context) {
	var userDTO dto.UserRequest

	user, msg := getMappingUser(userDTO, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	existIdentification := userService.IUser.IsDuplicateIdentification(user.Identification)

	if existIdentification {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.IDENTIFICATION_EXIST))
		return
	}
	existEmail, _ := userService.IUser.IsDuplicateEmail(user.Email)
	if existEmail {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.EMAIL_EXIST))
		return
	}
	user.Password = utilities.HashAndSalt([]byte(user.Password))

	filename, _ := UploadFile(context)

	user.Image = filename
	data, errs := userService.IUser.SetInsertUser(user)
	if errs != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(errs.Error()))
		return
	}
	result := getListsUserDto(data)
	context.JSON(http.StatusCreated, utilities.BuildCreatedResponse(result))
}

// update user
func (userService *userService) SetUpdateService(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	var userDTO dto.UserUpdateRequest

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	user, msg := getMappingUserUpdate(userDTO, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	existIdentification, _ := userService.IUser.IsDuplicateIdentificationById(user.Identification, uint(id))

	if existIdentification {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.IDENTIFICATION_EXIST))
		return
	}
	existEmail, _ := userService.IUser.IsDuplicateDiifEmailById(user.Email, uint(id))
	if existEmail {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.EMAIL_EXIST))
		return
	}
	findById, _ := userService.IUser.GetProfileUser(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}

	// filename, err := UploadFile(context)
	// if err != nil {
	// 	context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
	// 	return
	// }
	// if len(findById.Image) == 0 {
	// 	user.Image = filename
	// } else {
	// 	if filename != "" {
	// 		user.Image = filename
	// 	} else {
	// 		user.Image = findById.Image
	// 	}
	// }
	data, err := userService.IUser.SetEditUser(user, uint(id))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	result := getListsUserDto(data)
	context.JSON(http.StatusCreated, utilities.BuildUpdatedResponse(result))
}

func (userService *userService) SetUpdatePasswordService(context *gin.Context) {
	var userDTO dto.UserPasswordRequest
	user, msg := getMappingUserPassword(userDTO, context)
	if msg != "" {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	err := userService.IUser.SetChangePassword(user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildUpdatedResponse(nil))
}
func (userService *userService) SetRemoveService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	user, err := userService.IUser.GetProfileUser(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}

	ok, err := userService.IUser.SetRemoveUser(user.Id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.NOT_DELETED))
		return
	}
	if ok {
		result := getListsUserDto(user)
		context.JSON(http.StatusCreated, utilities.BuildRemovedResponse(result))
	}

}
func (userService *userService) GetProfileService(Id uint, context *gin.Context) {
	if Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	user, err := userService.IUser.GetProfileUser(Id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(user))
}

func (userService *userService) GetFindUserService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	data, err := userService.IUser.GetProfileUser(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	user := getListsUserDto(data)
	context.JSON(http.StatusOK, utilities.BuildResponse(user))
}
func (userService *userService) GetFindUserNameLastNameService(context *gin.Context) {
	var listUsrs []dto.UserResponse
	search := context.Param("search")
	data, err := userService.IUser.GetFindUserNameLastName(search)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, item := range data {
		user := getListsUserDto(item)
		listUsrs = append(listUsrs, user)
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(listUsrs))

}
func (userService *userService) GetAllService(context *gin.Context) {
	var usersLists []dto.UserResponse

	total := userService.IUser.GetCountUser()
	var limit int64 = 9
	page, begin := utilities.Pagination(context, int(limit))
	pages := (total / limit)
	if (total % limit) != 0 {
		pages++
	}

	users, err := userService.IUser.GetPaginationUsers(begin, int(limit))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, data := range users {
		user := getListsUserDto(data)
		usersLists = append(usersLists, user)
	}
	context.JSON(http.StatusOK, struct {
		Data  []dto.UserResponse `json:"data"`
		Total int64              `json:"total"`
		Page  int                `json:"page"`
		Pages int64              `json:"pages"`
		Limit int64              `json:"limit"`
	}{
		Data:  usersLists,
		Total: total,
		Page:  page,
		Pages: pages,
		Limit: limit,
	})
}

// method private
// mapping user
func getMappingUser(userDTO dto.UserRequest, context *gin.Context) (entity.User, string) {
	user := entity.User{}
	err := context.ShouldBindJSON(&userDTO)
	if err != nil {
		return user, utilities.GetMsgErrorRequired(err)
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		return user, err.Error()
	}
	return user, ""
}
func getMappingUserUpdate(userDTO dto.UserUpdateRequest, context *gin.Context) (entity.User, string) {
	user := entity.User{}
	err := context.ShouldBindJSON(&userDTO)
	if err != nil {
		return user, utilities.GetMsgErrorRequired(err)
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		return user, err.Error()
	}
	return user, ""
}

// mapping update password
func getMappingUserPassword(userDTO dto.UserPasswordRequest, context *gin.Context) (entity.User, string) {
	user := entity.User{}
	err := context.ShouldBindJSON(&userDTO)
	if err != nil {
		return user, utilities.GetMsgErrorRequired(err)
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		return user, err.Error()
	}
	return user, ""
}

func getListsUserDto(data entity.User) dto.UserResponse {
	return dto.UserResponse{
		Id:                 data.Id,
		Image:              data.Image,
		Name:               data.Name,
		LastName:           data.LastName,
		Identification:     data.Identification,
		TypeIdentification: data.TypeIdentification,
		Birthday:           data.Birthday,
		Birthplace:         data.Birthplace,
		Gender:             data.Gender,
		Rh:                 data.Rh,
		CivilStatus:        data.CivilStatus,
		Email:              data.Email,
		Direction:          data.Direction,
		CellPhone:          data.CellPhone,
		PhoneNumber:        data.PhoneNumber,
		Position:           data.Position,
		Occupation:         data.Occupation,
		School:             data.School,
		Grade:              data.Grade,
		HobbiesInterests:   data.HobbiesInterests,
		Allergies:          data.Allergies,
		BaptismWater:       data.BaptismWater,
		BaptismSpirit:      data.BaptismSpirit,
		YearConversion:     data.YearConversion,
		ChurchName:         data.Church.Name,
		ChurchId:           data.ChurchId,
		RolName:            data.Rol.Name,
		RolId:              data.RolId,
		CityName:           data.City.Name,
		CityId:             data.CityId,
		Active:             data.Active,
	}
}
