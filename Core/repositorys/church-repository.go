package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func GetChurchInstance() Interfaces.IChurch {
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
@param church is the Church, of type struct
*/
func (db *OpenConnections) SetCreateChurch(church entity.Church) (entity.Church, error) {
	db.mux.Lock()
	err := db.connection.Save(&church).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return church, err
}

/*
@param church is the Church, of type struct
*/
func (db *OpenConnections) SetUpdateChurch(church entity.Church, Id uint) (entity.Church, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", Id).Save(&church).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return church, err
}

/*
@param Id is the of Church, is of type uint
*/
func (db *OpenConnections) GetFindChurchById(Id uint) (entity.Church, error) {
	var church entity.Church
	db.mux.Lock()
	err := db.connection.Find(&church, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return church, err
}
func (db *OpenConnections) GetAllChurch() ([]entity.Church, error) {
	var churchs []entity.Church
	db.mux.Lock()
	err := db.connection.Find(&churchs).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return churchs, err
}

/*
@param church is the of Church, is of type struct
*/
func (db *OpenConnections) SetRemoveChurch(church entity.Church) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&church).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
