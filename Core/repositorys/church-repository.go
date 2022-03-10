package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type ChurchRepository interface {
	SetCreateChurch(church entity.Church) (entity.Church, error)
	SetRemoveChurch(church entity.Church) (bool, error)
	GetFindChurchById(Id uint) (entity.Church, error)
	GetAllChurch() ([]entity.Church, error)
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

var errChanChurch = make(chan error, constantvariables.CHAN_VALUE)

/*
@param church is the Church, of type struct
*/
func (db *churchConnection) SetCreateChurch(church entity.Church) (entity.Church, error) {

	go func() {
		err := db.connection.Save(&church).Error
		defer entity.Closedb()
		errChanChurch <- err
	}()
	err := <-errChanChurch
	return church, err
}

/*
@param Id is the of Church, is of type uint
*/
func (db *churchConnection) GetFindChurchById(Id uint) (entity.Church, error) {
	var church entity.Church

	go func() {
		err := db.connection.Find(&church, Id).Error
		defer entity.Closedb()
		errChanChurch <- err
	}()
	err := <-errChanChurch
	return church, err
}
func (db *churchConnection) GetAllChurch() ([]entity.Church, error) {
	var churchs []entity.Church

	go func() {
		err := db.connection.Find(&churchs).Error
		defer entity.Closedb()
		errChanChurch <- err
	}()
	err := <-errChanChurch
	return churchs, err
}

/*
@param church is the of Church, is of type struct
*/
func (db *churchConnection) SetRemoveChurch(church entity.Church) (bool, error) {

	go func() {
		err := db.connection.Delete(&church).Error
		defer entity.Closedb()
		errChanChurch <- err
	}()
	err := <-errChanChurch
	if err == nil {
		return true, err
	}
	return false, err
}
