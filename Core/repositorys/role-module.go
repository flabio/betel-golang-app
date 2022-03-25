package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type RoleModule interface {
	SetCreateRoleModule(roleModule entity.RoleModule) (entity.RoleModule, error)
	SetRemoveRoleModule(Id uint) (bool, error)
}

type roleModuleConnection struct {
	connection *gorm.DB
}

func NewRoleModuleRepository() RoleModule {
	var db *gorm.DB = entity.DatabaseConnection()
	return &roleModuleConnection{
		connection: db,
	}
}

/*
@param roleModule, is a struct of RoleModule
*/
func (db *roleModuleConnection) SetCreateRoleModule(roleModule entity.RoleModule) (entity.RoleModule, error) {
	err := db.connection.Save(&roleModule).Error
	defer entity.Closedb()
	return roleModule, err
}

/*
@param Id, is a struct of RoleModule
*/
func (db *roleModuleConnection) SetRemoveRoleModule(Id uint) (bool, error) {
	err := db.connection.Delete(&entity.RoleModule{}, Id).Error
	defer entity.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}
