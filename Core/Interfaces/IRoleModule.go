package Interfaces

import "bete/Core/entity"

type IRoleModule interface {
	SetCreateRoleModule(roleModule entity.RoleModule) (entity.RoleModule, error)
	SetDeleteRoleModule(Id uint) (bool, error)
}
