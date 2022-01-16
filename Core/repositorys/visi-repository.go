package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type VisitRepository interface {
	Create(visit entity.Visit) (entity.Visit, error)
	All() ([]entity.Visit, error)
	FindById(Id uint) (entity.Visit, error)
	AllVisitByUser(userId uint, subDetachmentId uint) ([]entity.Visit, error)
	Remove(Id uint) (bool, error)
}
type visitConnection struct {
	connection *gorm.DB
}

func NewVisitConnection() VisitRepository {
	return &visitConnection{
		connection: entity.DatabaseConnection(),
	}
}
func (db *visitConnection) Create(visit entity.Visit) (entity.Visit, error) {
	errChan := make(chan error, 1)
	go func() {
		err := db.connection.Save(&visit).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return visit, err
}
func (db *visitConnection) All() ([]entity.Visit, error) {
	var visit []entity.Visit
	errChan := make(chan error, 1)
	go func() {
		err := db.connection.Find(&visit).Error
		errChan <- err
		defer entity.Closedb()
	}()
	err := <-errChan
	return visit, err
}

func (db *visitConnection) AllVisitByUser(userId uint, subDetachmentId uint) ([]entity.Visit, error) {
	var visit []entity.Visit
	errChan := make(chan error, 1)
	go func() {
		err := db.connection.Where("userid=?", userId).Where("subdetachmentid=?", subDetachmentId).Find(&visit).Error
		errChan <- err
		defer entity.Closedb()
	}()

	err := <-errChan
	return visit, err
}
func (db *visitConnection) FindById(Id uint) (entity.Visit, error) {
	var visit entity.Visit
	errChan := make(chan error, 1)
	go func() {
		err := db.connection.Find(&visit, Id).Error
		errChan <- err
		defer entity.Closedb()
	}()

	err := <-errChan
	return visit, err
}

func (db *visitConnection) Remove(Id uint) (bool, error) {
	errChan := make(chan error, 1)
	go func() {
		err := db.connection.Delete(&entity.Visit{}, Id).Error
		errChan <- err
		defer entity.Closedb()
	}()
	err := <-errChan
	if err != nil {
		return true, err
	}
	return false, err
}
