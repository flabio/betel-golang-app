package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func NewModuleRepository() Interfaces.IModule {
	var (
		_OPEN *OpenConnections
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnections{

			connection: entity.Factory(constantvariables.OPTION_FACTORY_DB),
		}
	})
	return _OPEN
}

/*
@param module,is a struct of Module
*/
func (db *OpenConnections) SetCreateModule(module entity.Module) (entity.Module, error) {
	db.mux.Lock()
	err := db.connection.Save(&module).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return module, err
}

/*
@param rolemodule, is a struct of RoleModule
*/
func (db *OpenConnections) SetCreateModuleRole(rolemodule entity.RoleModule) (entity.RoleModule, error) {
	db.mux.Lock()
	err := db.connection.Save(&rolemodule).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return rolemodule, err
}
func (db *OpenConnections) GetAllModule() ([]entity.Module, error) {
	var modules []entity.Module
	db.mux.Lock()
	err := db.connection.Preload("RoleModule").Find(&modules).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return modules, err
}

/*
@param Id, is a uint of RoleModule
*/
func (db *OpenConnections) GetAllByRoleModule(Id uint) ([]entity.RoleModule, error) {
	var modules []entity.RoleModule
	db.mux.Lock()
	//err := db.connection.Preload("RoleModule").Joins("left join role_modules on role_modules.module_id = modules.id").Where("role_modules.rol_id", Id).Find(&modules).Error
	err := db.connection.Preload("Module").Where("rol_id", Id).Find(&modules).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return modules, err
}

/*
@param Id, is a uint of Module
*/
func (db *OpenConnections) GetFindModuleById(Id uint) (entity.Module, error) {
	var module entity.Module
	db.mux.Lock()
	err := db.connection.Find(&module, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return module, err
}

/*
@param Id, is a uint of RoleModule
*/
func (db *OpenConnections) GetFindRoleModuleById(Id uint) (entity.RoleModule, error) {
	var module entity.RoleModule
	db.mux.Lock()
	err := db.connection.Find(&module, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return module, err
}

/*
@param Id, is a uint of Module
*/
func (db *OpenConnections) SetRemoveModule(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&entity.Module{}, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}

/*
@param Id, is a uint of RoleModule
*/
func (db *OpenConnections) SetRemoveRoleModule(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&entity.RoleModule{}, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
