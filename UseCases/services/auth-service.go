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

		comparedPassword := comparePassword(v.Password, []byte(password))
		if comparedPassword {
			if v.Email == email && comparedPassword {
				return res
			}
		}
		return nil
	}
	return nil
}

func (authService *authService) CreateUser(user dto.UserDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
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

func (authService *authService) GetIdRol(id uint) uint {

	user, err := authService.IUser.GetProfileUser(id)
	if err != nil {

		return 0
	}
	return user.Roles.RolId
}
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	return true
}
