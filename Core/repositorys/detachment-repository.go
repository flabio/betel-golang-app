package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type DetachmentRepository interface {
	SetCreateDetachment(detachment entity.Detachment) (entity.Detachment, error)
	SetRemoveDetachment(detachment entity.Detachment) (entity.Detachment, error)
	GetFindDetachmentById(Id uint) (entity.Detachment, error)
	GetAllDetachment() ([]entity.Detachment, error)
}

type detachmentConnection struct {
	connection *gorm.DB
}

func NewDetachmentRepository() DetachmentRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &detachmentConnection{
		connection: db,
	}
}

var errChanDetachment = make(chan error, constantvariables.CHAN_VALUE)

func (db *detachmentConnection) SetCreateDetachment(res entity.Detachment) (entity.Detachment, error) {
	go func() {
		err := db.connection.Save(&res).Error
		defer entity.Closedb()
		errChanDetachment <- err
	}()
	err := <-errChanDetachment
	return res, err
}
func (db *detachmentConnection) SetRemoveDetachment(res entity.Detachment) (entity.Detachment, error) {
	go func() {
		err := db.connection.Delete(&res).Error
		defer entity.Closedb()
		errChanDetachment <- err
	}()
	err := <-errChanDetachment
	return res, err
}

func (db *detachmentConnection) GetFindDetachmentById(Id uint) (entity.Detachment, error) {
	var result entity.Detachment
	go func() {
		err := db.connection.Find(&result, Id).Error
		defer entity.Closedb()
		errChanDetachment <- err
	}()
	err := <-errChanDetachment
	return result, err
}
func (db *detachmentConnection) GetAllDetachment() ([]entity.Detachment, error) {
	var results []entity.Detachment

	go func() {
		err := db.connection.Find(&results).Error
		defer entity.Closedb()
		errChanDetachment <- err
	}()
	err := <-errChanDetachment
	return results, err
}
