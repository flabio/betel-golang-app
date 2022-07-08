package InterfacesService

import (
	"bete/Core/entity"
	"bete/UseCases/dto"
)

type IAuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserDTO) entity.User
	FindByEmail(email string) (entity.User, error)
	IsDuplicateEmail(email string) (bool, error)
	GetIdRol(id uint) uint
}
