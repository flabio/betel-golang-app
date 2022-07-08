package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func GetCityInstance() Interfaces.ICity {
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

func (db *OpenConnections) GetAllCity() ([]entity.City, error) {

	var citys []entity.City
	db.mux.Lock()
	err := db.connection.Find(&citys).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return citys, err
}
