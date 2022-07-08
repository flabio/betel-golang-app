package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func GetScoutParentInstance() Interfaces.IScoutParent {
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
@param parentScout, is a struct of ParentScout
*/
func (db *OpenConnections) SetCreateParentScouts(parentScout entity.ParentScout) (entity.ParentScout, error) {
	db.mux.Lock()
	err := db.connection.Save(&parentScout).Error
	defer entity.Closedb()
	defer db.mux.Unlock()

	return parentScout, err
}
