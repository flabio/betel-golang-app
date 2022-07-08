package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func NewWeeksDetachmentRepository() Interfaces.IWeeksDetachment {
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
@param Id, is a uint
*/
func (db *OpenConnections) GetFindByIdWeeks(Id uint) ([]entity.WeeksDetachment, error) {
	var weeksdetachment []entity.WeeksDetachment
	db.mux.Lock()
	err := db.connection.Order("id asc").
		Where("sub_detachment_id", Id).
		Find(&weeksdetachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return weeksdetachment, err
}
