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

/*
@param detachment,is a struct of Detachment
*/

func (db *detachmentConnection) SetCreateDetachment(detachment entity.Detachment) (entity.Detachment, error) {
	go func() {
		err := db.connection.Save(&detachment).Error
		defer entity.Closedb()
		errChanDetachment <- err
	}()
	err := <-errChanDetachment
	return detachment, err
}

/*
@param detachment,is a struct of Detachment
*/
func (db *detachmentConnection) SetRemoveDetachment(detachment entity.Detachment) (entity.Detachment, error) {
	go func() {
		err := db.connection.Delete(&detachment).Error
		defer entity.Closedb()
		errChanDetachment <- err
	}()
	err := <-errChanDetachment
	return detachment, err
}

/*
@param Id,is a uint of Detachment
*/
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
