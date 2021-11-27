package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type ModuleRepository interface {
	CreateModule(module entity.Module) (entity.Module, error)
	UpdateModule(module entity.Module) (entity.Module, error)
	AddModule(rolemodule entity.RoleModule) (entity.RoleModule, error)
	AllModule() ([]entity.Module, error)
	AllByRoleModule(Id uint) ([]entity.Module, error)
	FindModuleById(Id uint) (entity.Module, error)
	DeleteModule(Id uint) (bool, error)
	DeleteRoleModule(Id uint) (bool, error)
	FindRoleModuleById(Id uint) (entity.RoleModule, error)
}

type moduleConnection struct {
	connection *gorm.DB
}

func NewModuleRepository() ModuleRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &moduleConnection{
		connection: db,
	}
}
func (db *moduleConnection) CreateModule(module entity.Module) (entity.Module, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&module).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return module, err
}
func (db *moduleConnection) UpdateModule(module entity.Module) (entity.Module, error) {
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Save(&module).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return module, err
}

func (db *moduleConnection) AddModule(rolemodule entity.RoleModule) (entity.RoleModule, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&rolemodule).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return rolemodule, err
}
func (db *moduleConnection) AllModule() ([]entity.Module, error) {
	var modules []entity.Module
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Preload("RoleModule").Find(&modules).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return modules, err
}
func (db *moduleConnection) AllByRoleModule(Id uint) ([]entity.Module, error) {
	var modules []entity.Module
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Preload("RoleModule").Joins("left join role_modules on role_modules.module_id = modules.id").Where("role_modules.rol_id", Id).Find(&modules).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return modules, err
}

func (db *moduleConnection) FindModuleById(Id uint) (entity.Module, error) {
	var module entity.Module
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Find(&module, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return module, err
}

func (db *moduleConnection) FindRoleModuleById(Id uint) (entity.RoleModule, error) {
	var module entity.RoleModule
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Find(&module, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return module, err
}
func (db *moduleConnection) DeleteModule(Id uint) (bool, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Delete(&entity.Module{}, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *moduleConnection) DeleteRoleModule(Id uint) (bool, error) {
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
