package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func NewParentScoutRepository() Interfaces.IParentScout {
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
@param parent, is a struct of ParentScout
*/
func (db *OpenConnections) SetCreateParentScout(parent entity.ParentScout) (entity.ParentScout, error) {
	db.mux.Lock()
	err := db.connection.Save(&parent).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return parent, err
}

/*
@param Id, is a uint of ParentScout
*/
func (db *OpenConnections) SetRemoveParentScout(id int) (bool, error) {

	db.mux.Lock()
	err := db.connection.Delete(&id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
