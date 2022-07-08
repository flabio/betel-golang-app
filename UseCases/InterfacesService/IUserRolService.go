package InterfacesService

import (
	"bete/Core/entity"
	"bete/UseCases/dto"
)

//UserService is a contract.....
type IUserRolService interface {
	SetCreateUserRolService(role dto.RoleDTO) entity.Role
	GetAllUserRoleService() []entity.Role
}
