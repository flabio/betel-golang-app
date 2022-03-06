package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserRolService interface {
	SetCreateUserRolService(role dto.RoleDTO) entity.Role
	GetAllUserRoleService() []entity.Role
}

type userRolService struct {
	userRolRepository repositorys.UserRolRepository
}

//NewUserService creates a new instance of UserService
func NewUserRolService() UserRolService {
	userRepo := repositorys.NewUserRolRepository()
	return &userRolService{
		userRolRepository: userRepo,
	}
}

func (userService *userRolService) GetAllUserRoleService() []entity.Role {
	return userService.userRolRepository.GetAllUserRole()
}

// user and rol
func (userService *userRolService) SetCreateUserRolService(roleDTO dto.RoleDTO) entity.Role {
	role := entity.Role{}

	err := smapping.FillStruct(&role, smapping.MapFields(&roleDTO))
	if err != nil {
		checkError(err)
	}
	res := userService.userRolRepository.SetInsertUserRol(role)
	return res
}
