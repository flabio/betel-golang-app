package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"

	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

var wg sync.WaitGroup

//UserService is a contract.....
type UserService interface {
	Create(context *gin.Context)

	Update(context *gin.Context)
	UpdatePassword(context *gin.Context)
	All(context *gin.Context)
	ListUser(context *gin.Context)
	ListKingsScouts(context *gin.Context)
	Delete(context *gin.Context)
	Profile(userID uint, context *gin.Context)
	FindUser(context *gin.Context)
	FindUserNameLastName(context *gin.Context)
}

type userService struct {
	userRepository repositorys.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService() UserService {
	var userRepo repositorys.UserRepository = repositorys.NewUserRepository()
	return &userService{
		userRepository: userRepo,
	}
}

//List user
func (userService *userService) ListUser(context *gin.Context) {

	users, err := userService.userRepository.AllUser()

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "OK", users))
}

// list ListKingsScouts

func (userService *userService) ListKingsScouts(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	users, err := userService.userRepository.ListKingsScouts(uint(id))

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "OK", users))
}

//Create user
func (userService *userService) Create(context *gin.Context) {
	userChan := make(chan int)
	option := 0

	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.UserDTO

	context.ShouldBind(&userDTO)

	if validarUser(userDTO, userService, context, option) {
		return
	}

	smapping.FillStruct(&userToCreated, smapping.MapFields(&userDTO))

	userToCreated.Password = utilities.HashAndSalt([]byte(userToCreated.Password))

	go func(userChan chan<- int) {
		filename, _ := UploadFile(context)

		userToCreated.Image = filename
		createdUser, errs := userService.userRepository.InsertUser(userToCreated)
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

		roleToCreated.RolId = userDTO.RolId
		roleToCreated.UserId = uint(user_id)

		err := userService.userRepository.InsertRole(roleToCreated)
		if err != nil {
			log.Println(err)
			result, err := userService.userRepository.DeleteUser(uint(user_id))
			if err != nil {
				validadErrorRemove(result, context)
				return
			}
			validadErrors(err, context)
			return
		}
	}
}

//update user
func (userService *userService) Update(context *gin.Context) {
	userToCreated := entity.User{}
	roleToCreated := entity.Role{}

	var userDTO dto.UserDTO

	context.ShouldBind(&userDTO)

	if validarUser(userDTO, userService, context, 1) {
		return
	}

	smapping.FillStruct(&userToCreated, smapping.MapFields(&userDTO))

	roleToCreated.RolId = userDTO.RolId
	roleToCreated.UserId = userDTO.Id

	wg.Add(1)
	go goRunitaUpdateRole(userService, roleToCreated)
	wg.Wait()

	findById, _ := userService.userRepository.ProfileUser(uint(userDTO.Id))
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
	u, err := userService.userRepository.EditUser(userToCreated)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(u)
	context.JSON(http.StatusOK, res)

}
func (userService *userService) UpdatePassword(context *gin.Context) {
	user := entity.User{}
	var userDTO dto.UserPasswordDTO

	errDTO := context.ShouldBind(&userDTO)
	if errDTO != nil {
		validadErrors(errDTO, context)
		return
	}
	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	checkError(err)
	errp := userService.userRepository.ChangePassword(user)
	if errp != nil {
		validadErrors(errp, context)
		return
	}
	res := utilities.BuildUpdatePasswordResponse()
	context.JSON(http.StatusOK, res)
}
func (userService *userService) Delete(context *gin.Context) {
	chanels := make(chan bool)
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	user, errprofile := userService.userRepository.ProfileUser(uint(id))
	if errprofile != nil {
		validadErrorById(context)
		return
	}

	go func() {
		err := userService.userRepository.DeleteRoleUser(user.Id)
		if err != nil {
			chanels <- false
			return
		}
		chanels <- true
		close(chanels)
	}()
	if <-chanels {

		result, err := userService.userRepository.DeleteUser(user.Id)
		if err != nil {
			validadErrorRemove(result, context)
			return
		}
		res := utilities.BuildDeteleteResponse(true, user)
		context.JSON(http.StatusOK, res)
		return
	}
	res := utilities.BuildNotFoudResponse()
	context.JSON(http.StatusBadRequest, res)
}
func (userService *userService) Profile(Id uint, context *gin.Context) {

	user, err := userService.userRepository.ProfileUser(Id)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}

func (userService *userService) FindUser(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	user, err := userService.userRepository.ProfileUser(uint(id))
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}
func (userService *userService) FindUserNameLastName(context *gin.Context) {

	search := context.Param("search")

	users, err := userService.userRepository.FindUserNameLastName(search)

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "OK", users))
	return

}
func (userService *userService) All(context *gin.Context) {
	total := userService.userRepository.CountUser()
	var limit int64 = 9
	page, begin := utilities.Pagination(context, int(limit))
	pages := (total / limit)
	if (total % limit) != 0 {
		pages++
	}
	fmt.Printf("Current Page: %d, Begin: %d\n", page, begin)
	users, err := userService.userRepository.PaginationUsers(begin, int(limit))

	if err != nil {
		validadErrors(err, context)
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
	err := userService.userRepository.InsertRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
	}
}

//goRunitaUpdateRole
func goRunitaUpdateRole(userService *userService, roleToCreated entity.Role) {
	wg.Done()
	role, err := userService.userRepository.EditRole(roleToCreated)
	if err != nil {
		log.Println(err)
		checkError(err)
	}

	if role.Id == 0 {
		userService.userRepository.InsertRole(roleToCreated)
	}
}

//validarUser
func validarUser(u dto.UserDTO, userService *userService, context *gin.Context, option int) bool {
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
	if len(u.CivilStatus) == 0 {
		validadRequiredMsg(msg.RequiredCivilStatus(), context)
		return true
	}
	if len(u.Email) == 0 {
		validadRequiredMsg(msg.RequiredEmail(), context)
		return true
	}

	if u.RolId == 0 {
		validadRequiredMsg(msg.RequiredRol(), context)
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
		existEmail, _ := userService.userRepository.IsDuplicateEmail(u.Email)
		existsIdentification := userService.userRepository.IsDuplicateIdentificatio(u.Identification)
		if existsIdentification {
			validadRequiredMsg(msg.RequiredExistIdentification(), context)
			return true
		}
		if existEmail {
			validadRequiredMsg(msg.RequiredExistEmail(), context)
			return true
		}
		if u.Password == "" {
			validadRequiredMsg(msg.RequiredPassword(), context)
			return true
		}
		if u.Password != u.ConfirmPassword {
			validadRequiredMsg(msg.RequiredPasswordConfirm(), context)
			return true
		}
	}
	return false
}
