package Interfaces

import (
	"bete/Core/entity"
)

type IRol interface {
	SetCreateRol(rol entity.Rol) (entity.Rol, error)
	SetUpdateRol(rol entity.Rol, Id uint) (entity.Rol, error)
	GetAllRol() ([]entity.Rol, error)
	GetAllGroupRol() ([]entity.Rol, error)
	GetRolsModule() ([]entity.RoleModule, error)
	SetRemoveRol(rol entity.Rol) (bool, error)
	GetFindRolById(Id uint) (entity.Rol, error)
}
