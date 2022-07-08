package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
)

type userRolService struct {
	IUserRol Interfaces.IUserRol
}

//NewUserService creates a new instance of UserService
func NewUserRolService() InterfacesService.IUserRolService {
	return &userRolService{
		IUserRol: repositorys.NewUserRolRepository(),
	}
}

func (userService *userRolService) GetAllUserRoleService() []entity.Role {
	return userService.IUserRol.GetAllUserRole()
}

// user and rol
func (userService *userRolService) SetCreateUserRolService(roleDTO dto.RoleDTO) entity.Role {
	role := entity.Role{}

	err := smapping.FillStruct(&role, smapping.MapFields(&roleDTO))
	if err != nil {
		checkError(err)
	}
	res := userService.IUserRol.SetInsertUserRol(role)
	return res
}
