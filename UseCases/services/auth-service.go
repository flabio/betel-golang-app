package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//authService is a contract about something that this service can do

type authService struct {
	IUser Interfaces.IUser
}

//NewauthService creates a new instance of authService
func NewAuthService() InterfacesService.IAuthService {

	return &authService{
		IUser: repositorys.NewUserRepository(),
	}
}

func (authService *authService) VerifyCredential(email string, password string) interface{} {
	res := authService.IUser.VerifyCredential(email, password)

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

	res, err := authService.IUser.SetInsertUser(userToCreate)

	return res
}

func (authService *authService) FindByEmail(email string) (entity.User, error) {
	return authService.IUser.GetFindByEmail(email)
}

func (authService *authService) IsDuplicateEmail(email string) (bool, error) {
	return authService.IUser.IsDuplicateEmail(email)

}

/*
@param Id is of type uint
*/
func (authService *authService) GetIdRol(Id uint) uint {

	user, err := authService.IUser.GetProfileUser(Id)
	if err != nil {

		return 0
	}
	return user.Roles.RolId
}
func comparePassword(HashedPwd string, PlainPassword []byte) bool {
	byteHash := []byte(HashedPwd)
	bcrypt.CompareHashAndPassword(byteHash, PlainPassword)

	return true
}
