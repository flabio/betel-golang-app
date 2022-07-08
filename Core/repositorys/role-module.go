package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func GetRoleModuleInstance() Interfaces.IRoleModule {
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
@param roleModule, is a struct of RoleModule
*/
func (db *OpenConnections) SetCreateRoleModule(roleModule entity.RoleModule) (entity.RoleModule, error) {
	db.mux.Lock()
	err := db.connection.Save(&roleModule).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return roleModule, err
}

/*
@param Id, is a struct of RoleModule
*/
func (db *OpenConnections) SetDeleteRoleModule(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&entity.RoleModule{}, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
