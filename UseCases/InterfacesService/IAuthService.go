package InterfacesService

import (
	"bete/Core/entity"
)

type IAuthService interface {
	VerifyCredential(email string, password string) interface{}
	//CreateUser(user dto.UserRequest, context *gin.Context) entity.User
	FindByEmail(email string) (entity.User, error)
	IsDuplicateEmail(email string) (bool, error)
	GetIdRol(id uint) uint
}
