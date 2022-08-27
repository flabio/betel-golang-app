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
	var usersLists []dto.UserListDTO
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
	var userDTO dto.UserDTO

	user, err := getMappingUser(userDTO, context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	user.Password = utilities.HashAndSalt([]byte(user.Password))

	filename, _ := UploadFile(context)

	user.Image = filename
	createdUser, errs := userService.IUser.SetInsertUser(user)
	if errs != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(errs.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildCreatedResponse(createdUser))
}

// update user
func (userService *userService) SetUpdateService(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	var userDTO dto.UserUpdateDTO

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	user, err := getMappingUserUpdate(userDTO, context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	findById, _ := userService.IUser.GetProfileUser(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	filename, err := UploadFile(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if len(findById.Image) == 0 {
		user.Image = filename
	} else {
		if filename != "" {
			user.Image = filename
		} else {
			user.Image = findById.Image
		}
	}
	u, err := userService.IUser.SetEditUser(user, uint(id))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildUpdatedResponse(u))
}

func (userService *userService) SetUpdatePasswordService(context *gin.Context) {
	var userDTO dto.UserPasswordDTO
	user, err := getMappingUserPassword(userDTO, context)
	if err != nil {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	err = userService.IUser.SetChangePassword(user)
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
		context.JSON(http.StatusCreated, utilities.BuildRemovedResponse(user))
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
	user, err := userService.IUser.GetProfileUser(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(user))
}
func (userService *userService) GetFindUserNameLastNameService(context *gin.Context) {

	search := context.Param("search")

	users, err := userService.IUser.GetFindUserNameLastName(search)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(users))

}
func (userService *userService) GetAllService(context *gin.Context) {
	var usersLists []dto.UserListDTO

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
		Data  []dto.UserListDTO `json:"data"`
		Total int64             `json:"total"`
		Page  int               `json:"page"`
		Pages int64             `json:"pages"`
		Limit int64             `json:"limit"`
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
func getMappingUser(userDTO dto.UserDTO, context *gin.Context) (entity.User, error) {
	user := entity.User{}
	err := context.ShouldBindJSON(&userDTO)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return user, err
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return user, err
	}
	return user, nil
}
func getMappingUserUpdate(userDTO dto.UserUpdateDTO, context *gin.Context) (entity.User, error) {
	user := entity.User{}
	err := context.ShouldBindJSON(&userDTO)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return user, err
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return user, err
	}
	return user, nil
}

// mapping update password
func getMappingUserPassword(userDTO dto.UserPasswordDTO, context *gin.Context) (entity.User, error) {
	user := entity.User{}
	err := context.ShouldBindJSON(&userDTO)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return user, err
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return user, err
	}
	return user, nil
}

func getListsUserDto(data entity.User) dto.UserListDTO {
	return dto.UserListDTO{
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
