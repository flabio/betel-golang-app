package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

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

var errChanRoleModule = make(chan error, constantvariables.CHAN_VALUE)

func (db *roleModuleConnection) SetCreateRoleModule(roleModule entity.RoleModule) (entity.RoleModule, error) {
	go func() {
		err := db.connection.Save(&roleModule).Error
		defer entity.Closedb()
		errChanRoleModule <- err
	}()
	err := <-errChanRoleModule
	return roleModule, err
}
func (db *roleModuleConnection) SetRemoveRoleModule(Id uint) (bool, error) {
	go func() {
		err := db.connection.Delete(&entity.RoleModule{}, Id).Error
		defer entity.Closedb()
		errChanRoleModule <- err
	}()
	err := <-errChanRoleModule
	if err == nil {
		return true, err
	}
	return false, err
}
