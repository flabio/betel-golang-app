package Interfaces

import "bete/Core/entity"

type IModule interface {
	SetCreateModule(module entity.Module) (entity.Module, error)
	SetUpdateModule(module entity.Module, Id uint) (entity.Module, error)

	SetCreateModuleRole(rolemodule entity.RoleModule) (entity.RoleModule, error)
	GetAllModule() ([]entity.Module, error)
	GetAllByRoleModule(Id uint) ([]entity.RoleModule, error)
	GetFindModuleById(Id uint) (entity.Module, error)
	SetRemoveModule(Id uint) (bool, error)
	SetRemoveRoleModule(Id uint) (bool, error)
	GetFindRoleModuleById(Id uint) (entity.RoleModule, error)
}
