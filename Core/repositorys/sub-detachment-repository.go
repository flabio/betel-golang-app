package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type SubDetachmentRepository interface {
	SetCreateSubDetachment(subdetachment entity.SubDetachment) (entity.SubDetachment, error)
	SetRemoveSubDetachment(Id uint) (bool, error)
	GetFindByIdSubDetachment(Id uint) (entity.SubDetachment, error)
	GetFindByIdDetachmentSubDetachment(Id uint) ([]entity.SubDetachment, error)
	GetAllSubDetachment() ([]entity.SubDetachment, error)
}

type subConnection struct {
	connection *gorm.DB
}

func NewSubDetachmentRepository() SubDetachmentRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &subConnection{
		connection: db,
	}
}

var errChanSubdetachment = make(chan error, constantvariables.CHAN_VALUE)

func (db *subConnection) SetCreateSubDetachment(subdetachment entity.SubDetachment) (entity.SubDetachment, error) {
	go func() {
		err := db.connection.Save(&subdetachment).Error
		defer entity.Closedb()
		errChanSubdetachment <- err
	}()
	return subdetachment, <-errChan
}

func (db *subConnection) SetRemoveSubDetachment(Id uint) (bool, error) {
	go func() {
		err := db.connection.Delete(&entity.SubDetachment{}, Id).Error
		defer entity.Closedb()
		errChanSubdetachment <- err
	}()
	err := <-errChanSubdetachment
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
func (db *subConnection) GetFindByIdSubDetachment(Id uint) (entity.SubDetachment, error) {
	var subdetachment entity.SubDetachment
	go func() {
		err := db.connection.Find(&subdetachment, Id).Error
		defer entity.Closedb()
		errChanSubdetachment <- err
	}()
	err := <-errChanSubdetachment
	return subdetachment, err
}
func (db *subConnection) GetAllSubDetachment() ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	go func() {
		err := db.connection.Order("id desc").Preload("Detachment").Find(&subdetachment).Error
		defer entity.Closedb()
		errChanSubdetachment <- err
	}()
	err := <-errChanSubdetachment
	return subdetachment, err
}
func (db *subConnection) GetFindByIdDetachmentSubDetachment(Id uint) ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	go func() {
		err := db.connection.Order("id desc").Preload("Detachment").Where("detachment_id", Id).Find(&subdetachment).Error
		defer entity.Closedb()
		errChanSubdetachment <- err
	}()
	err := <-errChanSubdetachment
	return subdetachment, err
}
