package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type ModuleRepository interface {
	SetCreateModule(module entity.Module) (entity.Module, error)
	SetCreateModuleRole(rolemodule entity.RoleModule) (entity.RoleModule, error)
	GetAllModule() ([]entity.Module, error)
	GetAllByRoleModule(Id uint) ([]entity.RoleModule, error)
	GetFindModuleById(Id uint) (entity.Module, error)
	SetRemoveModule(Id uint) (bool, error)
	SetRemoveRoleModule(Id uint) (bool, error)
	GetFindRoleModuleById(Id uint) (entity.RoleModule, error)
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

var errChanModule = make(chan error, constantvariables.CHAN_VALUE)

/*
@param module,is a struct of Module
*/
func (db *moduleConnection) SetCreateModule(module entity.Module) (entity.Module, error) {
	go func() {
		err := db.connection.Save(&module).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	return module, err
}

/*
@param rolemodule, is a struct of RoleModule
*/
func (db *moduleConnection) SetCreateModuleRole(rolemodule entity.RoleModule) (entity.RoleModule, error) {

	go func() {
		err := db.connection.Save(&rolemodule).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	return rolemodule, err
}
func (db *moduleConnection) GetAllModule() ([]entity.Module, error) {
	var modules []entity.Module
	go func() {
		err := db.connection.Preload("RoleModule").Find(&modules).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	return modules, err
}

/*
@param Id, is a uint of RoleModule
*/
func (db *moduleConnection) GetAllByRoleModule(Id uint) ([]entity.RoleModule, error) {
	var modules []entity.RoleModule
	go func() {
		//err := db.connection.Preload("RoleModule").Joins("left join role_modules on role_modules.module_id = modules.id").Where("role_modules.rol_id", Id).Find(&modules).Error
		err := db.connection.Preload("Module").Where("rol_id", Id).Find(&modules).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	return modules, err
}

/*
@param Id, is a uint of Module
*/
func (db *moduleConnection) GetFindModuleById(Id uint) (entity.Module, error) {
	var module entity.Module
	go func() {
		err := db.connection.Find(&module, Id).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	return module, err
}

/*
@param Id, is a uint of RoleModule
*/
func (db *moduleConnection) GetFindRoleModuleById(Id uint) (entity.RoleModule, error) {
	var module entity.RoleModule
	go func() {
		err := db.connection.Find(&module, Id).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	return module, err
}

/*
@param Id, is a uint of Module
*/
func (db *moduleConnection) SetRemoveModule(Id uint) (bool, error) {
	go func() {
		err := db.connection.Delete(&entity.Module{}, Id).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	if err == nil {
		return true, err
	}
	return false, err
}

/*
@param Id, is a uint of RoleModule
*/
func (db *moduleConnection) SetRemoveRoleModule(Id uint) (bool, error) {
	go func() {
		err := db.connection.Delete(&entity.RoleModule{}, Id).Error
		defer entity.Closedb()
		errChanModule <- err
	}()
	err := <-errChanModule
	if err == nil {
		return true, err
	}
	return false, err
}
