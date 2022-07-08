package Interfaces

import "bete/Core/entity"

type IUserRol interface {
	SetInsertUserRol(role entity.Role) entity.Role
	GetAllUserRole() []entity.Role
}
