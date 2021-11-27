package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type RoleModule interface {
	Create(roleModule entity.RoleModule) (entity.RoleModule, error)
	Remove(Id uint) (bool, error)
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

func (db *roleModuleConnection) Create(roleModule entity.RoleModule) (entity.RoleModule, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&roleModule).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return roleModule, err
}
func (db *roleModuleConnection) Remove(Id uint) (bool, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Delete(&entity.RoleModule{}, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return true, err
	}
	return false, err
}
