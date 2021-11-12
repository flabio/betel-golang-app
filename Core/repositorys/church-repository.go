package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type ChurchRepository interface {
	CreateChurch(church entity.Church) (entity.Church, error)
	UpdateChurch(church entity.Church) (entity.Church, error)
	DeleteChurch(church entity.Church) (bool, error)
	FindChurchById(Id uint) (entity.Church, error)
	AllChurch() ([]entity.Church, error)
}

type churchConnection struct {
	connection *gorm.DB
}

func NewChurchRepository() ChurchRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &churchConnection{
		connection: db,
	}
}

func (db *churchConnection) CreateChurch(church entity.Church) (entity.Church, error) {
	var errChan=make(chan error,1)
	go func (db *churchConnection){
		err := db.connection.Save(&church).Error
		defer entity.Closedb()
		errChan<-err
	}(db)
	err := <-errChan
	return church, err
}
func (db *churchConnection) UpdateChurch(church entity.Church) (entity.Church, error) {
	var errChan=make(chan error,1)
	go func (db *churchConnection){
		err := db.connection.Save(&church).Error
		defer entity.Closedb()
		errChan<-err
	}(db)
	err := <-errChan
	return church, err
}
func (db *churchConnection) FindChurchById(Id uint) (entity.Church, error) {
	var church entity.Church
	var errChan=make(chan error,1)
	go func (db *churchConnection){
		err := db.connection.Find(&church, Id).Error
		defer entity.Closedb()
		errChan<-err
	}(db)
	err := <-errChan
	return church, err
}
func (db *churchConnection) AllChurch() ([]entity.Church, error) {
	var churchs []entity.Church
	var errChan=make(chan error,1)
	go func (db *churchConnection){
		err := db.connection.Find(&churchs).Error
		defer entity.Closedb()
		errChan<-err
	}(db)
	err := <-errChan
	return churchs, err
}
func (db *churchConnection) DeleteChurch(church entity.Church) (bool, error) {

	var errChan=make(chan error,1)
	go func (db *churchConnection){
		err := db.connection.Delete(&church).Error
		defer entity.Closedb()
		errChan<-err
	}(db)
	err := <-errChan
	if err == nil {
		return true, err
	}
	return false, err
}
