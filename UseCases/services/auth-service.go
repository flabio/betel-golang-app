package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//authService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(Emmail string, Password string) interface{}
	CreateUser(User dto.UserDTO) entity.User
	FindByEmail(Email string) (entity.User, error)
	IsDuplicateEmail(Email string) (bool, error)
	GetIdRol(Id uint) uint
}

type authService struct {
	userRepository repositorys.UserRepository
}

//NewauthService creates a new instance of authService
func NewAuthService() AuthService {

	var userRepo repositorys.UserRepository = repositorys.NewUserRepository()
	return &authService{
		userRepository: userRepo,
	}
}

/*
@param Email is of type string
@param Password is of type string
*/
func (authService *authService) VerifyCredential(Email string, Password string) interface{} {
	res := authService.userRepository.VerifyCredential(Email, Password)

	if v, ok := res.(entity.User); ok {
		if v.Email == Email {
			return res
		}
		//comparedPassword := comparePassword(v.Password, []byte(Password))
		/*if comparedPassword {
			if v.Email == Email && comparedPassword {
				return res
			}
		}*/
		return nil
	}
	return nil
}

/*
@param User is of type struct

*/
func (authService *authService) CreateUser(User dto.UserDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&User))
	checkError(err)

	res, err := authService.userRepository.SetInsertUser(userToCreate)

	return res
}

/*
@param Email is of type string
*/
func (authService *authService) FindByEmail(Email string) (entity.User, error) {
	return authService.userRepository.GetFindByEmail(Email)
}

/*
@param Email is of type string
*/
func (authService *authService) IsDuplicateEmail(Email string) (bool, error) {
	return authService.userRepository.IsDuplicateEmail(Email)

}

/*
@param Id is of type uint
*/
func (authService *authService) GetIdRol(Id uint) uint {

	user, _ := authService.userRepository.GetProfileUser(Id)
	return user.Roles.RolId
}
func comparePassword(HashedPwd string, PlainPassword []byte) bool {
	byteHash := []byte(HashedPwd)
	bcrypt.CompareHashAndPassword(byteHash, PlainPassword)

	return true
}
