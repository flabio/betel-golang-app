package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func GetRolInstance() Interfaces.IRol {

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
@param rol, is a struct of Rol
*/
func (db *OpenConnections) SetCreateRol(rol entity.Rol) (entity.Rol, error) {

	db.mux.Lock()
	err := db.connection.Save(&rol).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return rol, err
}

func (db *OpenConnections) SetUpdateRol(rol entity.Rol, Id uint) (entity.Rol, error) {

	db.mux.Lock()
	err := db.connection.Where("id=?", Id).Updates(&rol).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return rol, err
}

/*
@param rol, is a struct of Rol
*/
func (db *OpenConnections) SetRemoveRol(rol entity.Rol) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&rol).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}

/*
@param Id, is a uint of Rol
*/
func (db *OpenConnections) GetFindRolById(Id uint) (entity.Rol, error) {

	var rol entity.Rol
	db.mux.Lock()
	err := db.connection.Find(&rol, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return rol, err
}

func (db *OpenConnections) GetAllRol() ([]entity.Rol, error) {

	var rols []entity.Rol
	db.mux.Lock()
	err := db.connection.Find(&rols).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return rols, err
}

func (db *OpenConnections) GetAllGroupRol() ([]entity.Rol, error) {
	var rols []entity.Rol
	db.mux.Lock()
	err := db.connection.Where("id IN ?", []int{
		constantvariables.NAVIGANTORS_ROL,
		constantvariables.PIONEERS_ROL,
		constantvariables.PATH_FOLLOWERS_ROL,
		constantvariables.SCOUTS_ROL}).
		Find(&rols).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return rols, err
}
func (db *OpenConnections) GetRolsModule() ([]entity.RoleModule, error) {

	var roleModule []entity.RoleModule
	db.mux.Lock()
	err := db.connection.Preload("Role.Rol").
		Preload("Module").
		Find(&roleModule).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return roleModule, err
}
