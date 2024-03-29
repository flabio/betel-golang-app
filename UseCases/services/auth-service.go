package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/utilities"
)

//authService is a contract about something that this service can do

type authService struct {
	IUser Interfaces.IUser
}

// NewauthService creates a new instance of authService
func NewAuthService() InterfacesService.IAuthService {

	return &authService{
		IUser: repositorys.NewUserRepository(),
	}
}

func (authService *authService) VerifyCredential(email string, password string) interface{} {
	res := authService.IUser.VerifyCredential(email, password)

	if v, ok := res.(entity.User); ok {
		if v.Email == email {
			return v
		}
		comparedPassword := utilities.ComparePassword(v.Password, []byte(password))
		if comparedPassword {
			if v.Email == email && comparedPassword {
				return res
			}
		}
		return nil
	}
	return nil
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
	return user.RolId
}
