package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type DetachmentRepository interface {
	CreateDetachment(detachment entity.Detachment) (entity.Detachment, error)
	UpdateDetachment(detachment entity.Detachment) (entity.Detachment, error)
	DeleteDetachment(detachment entity.Detachment) (entity.Detachment, error)
	FindDetachmentById(Id uint) (entity.Detachment, error)
	AllDetachment() ([]entity.Detachment, error)
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

func (db *detachmentConnection) CreateDetachment(res entity.Detachment) (entity.Detachment, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&res).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return res, err
}
func (db *detachmentConnection) UpdateDetachment(res entity.Detachment) (entity.Detachment, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&res).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return res, err
}

func (db *detachmentConnection) DeleteDetachment(res entity.Detachment) (entity.Detachment, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Delete(&res).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return res, err
}

func (db *detachmentConnection) FindDetachmentById(Id uint) (entity.Detachment, error) {
	var result entity.Detachment
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Find(&result, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return result, err
}
func (db *detachmentConnection) AllDetachment() ([]entity.Detachment, error) {
	var results []entity.Detachment

	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Find(&results).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return results, err
}
