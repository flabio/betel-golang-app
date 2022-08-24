package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func GetDetachmentInstance() Interfaces.IDetachment {
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
@param detachment,is a struct of Detachment
*/

func (db *OpenConnections) SetCreateDetachment(detachment entity.Detachment) (entity.Detachment, error) {
	db.mux.Lock()
	err := db.connection.Save(&detachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return detachment, err
}

func (db *OpenConnections) SetUpdateDetachment(detachment entity.Detachment, Id uint) (entity.Detachment, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", Id).Save(&detachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return detachment, err
}

/*
@param detachment,is a struct of Detachment
*/
func (db *OpenConnections) SetRemoveDetachment(detachment entity.Detachment) (entity.Detachment, error) {
	db.mux.Lock()
	err := db.connection.Delete(&detachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return detachment, err
}

/*
@param Id,is a uint of Detachment
*/
func (db *OpenConnections) GetFindDetachmentById(Id uint) (entity.Detachment, error) {
	var result entity.Detachment
	db.mux.Lock()
	err := db.connection.Find(&result, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return result, err
}
func (db *OpenConnections) GetAllDetachment() ([]entity.Detachment, error) {
	var results []entity.Detachment
	db.mux.Lock()
	err := db.connection.Preload("Church").Find(&results).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return results, err
}
