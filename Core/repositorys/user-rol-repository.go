package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

//NewUserRepository is creates a new instance of UserRepository

func NewUserRolRepository() Interfaces.IUserRol {
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
@param rol, is a struct of Role
*/
func (db *OpenConnections) SetInsertUserRol(rol entity.Role) entity.Role {
	db.mux.Lock()
	db.connection.Save(&rol)
	defer entity.Closedb()
	defer db.mux.Unlock()
	return rol
}

func (db *OpenConnections) GetAllUserRole() []entity.Role {
	var role []entity.Role
	db.mux.Lock()
	db.connection.Joins("User").Joins("Rol").Find(&role)
	defer entity.Closedb()
	defer db.mux.Unlock()
	return role
}
