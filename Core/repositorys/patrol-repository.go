package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func NewPatrolRepository() Interfaces.IPatrol {
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
@param patrol, is a struct of Patrol
*/
func (db *OpenConnections) SetCreatePatrol(patrol entity.Patrol) (entity.Patrol, error) {
	db.mux.Lock()
	err := db.connection.Save(&patrol).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return patrol, err
}

/*
@param Id, is a uint of Patrol
*/
func (db *OpenConnections) SetRemovePatrol(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&entity.Patrol{}, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return false, err
	}
	return true, err
}

/*
@param Id, is a uint of Patrol
*/
func (db *OpenConnections) GetFindByIdPatrol(Id uint) (entity.Patrol, error) {
	var patrol entity.Patrol
	db.mux.Lock()
	err := db.connection.Find(&patrol, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return patrol, err
}

func (db *OpenConnections) GetAllPatrol() ([]entity.Patrol, error) {
	var patrol []entity.Patrol
	db.mux.Lock()
	err := db.connection.Order("id desc").Preload("SubDetachment").Find(&patrol).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return patrol, err
}
