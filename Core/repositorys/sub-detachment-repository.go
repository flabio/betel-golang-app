package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"

	"gorm.io/gorm"
)

type subConnection struct {
	connection *gorm.DB
	mux        sync.Mutex
}

func NewSubDetachmentRepository() Interfaces.ISubDetachment {
	var (
		_ONCE           sync.Once
		_SUB_DETACHMENT *subConnection
	)
	_ONCE.Do(func() {
		_SUB_DETACHMENT = &subConnection{
			connection: entity.Factory(constantvariables.OPTION_FACTORY_DB),
		}
	})
	return _SUB_DETACHMENT
}

/*
@param subdetachment, is a struct of SubDetachment
*/
func (db *subConnection) SetCreateSubDetachment(subdetachment entity.SubDetachment) (entity.SubDetachment, error) {
	db.mux.Lock()
	err := db.connection.Save(&subdetachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return subdetachment, err
}

/*
@param subdetachment, is a struct of SubDetachment
*/
func (db *subConnection) SetUpdateSubDetachment(subdetachment entity.SubDetachment, Id uint) (entity.SubDetachment, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", Id).Save(&subdetachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return subdetachment, err
}

/*
@param Id, is a uint of SubDetachment
*/
func (db *subConnection) SetRemoveSubDetachment(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&entity.SubDetachment{}, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return false, err
	}
	return true, err
}

/*
@param Id, is a uint of SubDetachment
*/
func (db *subConnection) GetFindByIdSubDetachment(Id uint) (entity.SubDetachment, error) {
	var subdetachment entity.SubDetachment
	db.mux.Lock()
	err := db.connection.Preload("Detachment").Find(&subdetachment, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return subdetachment, err
}
func (db *subConnection) GetAllSubDetachment() ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	db.mux.Lock()
	err := db.connection.Order("id desc").Preload("Detachment").Find(&subdetachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return subdetachment, err
}

/*
@param Id, is a uint of SubDetachment
*/
func (db *subConnection) GetFindByIdDetachmentSubDetachment(Id uint) ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	db.mux.Lock()
	err := db.connection.Order("id desc").Preload("Detachment").Where("detachment_id", Id).Find(&subdetachment).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return subdetachment, err
}
