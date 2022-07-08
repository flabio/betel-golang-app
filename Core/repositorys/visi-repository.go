package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func GetVisitInstance() Interfaces.IVisit {

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
@param visit, is a struct of Visit
*/
func (db *OpenConnections) SetCreateVisit(visit entity.Visit) (entity.Visit, error) {
	db.mux.Lock()
	err := db.connection.Save(&visit).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return visit, err
}
func (db *OpenConnections) GetAllVisit() ([]entity.Visit, error) {
	var visit []entity.Visit
	db.mux.Lock()
	err := db.connection.Find(&visit).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return visit, err
}

/*
@param userId, is a uint
@param subDetachmentId, is a uint
*/
func (db *OpenConnections) GetAllVisitByUserVisit(userId uint, subDetachmentId uint) ([]entity.Visit, error) {
	var visit []entity.Visit
	db.mux.Lock()
	err := db.connection.Where("userid=?", userId).
		Where("subdetachmentid=?", subDetachmentId).
		Find(&visit).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return visit, err
}

/*
@param Id, is a uint
*/
func (db *OpenConnections) GetFindByIdVisit(Id uint) (entity.Visit, error) {
	var visit entity.Visit
	db.mux.Lock()
	err := db.connection.Find(&visit, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return visit, err
}

/*
@param Id, is a uint
*/
func (db *OpenConnections) SetRemoveVisit(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&entity.Visit{}, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return true, err
	}
	return false, err
}
