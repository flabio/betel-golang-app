package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type VisitRepository interface {
	SetCreateVisit(visit entity.Visit) (entity.Visit, error)
	GetAllVisit() ([]entity.Visit, error)
	GetFindByIdVisit(Id uint) (entity.Visit, error)
	GetAllVisitByUserVisit(userId uint, subDetachmentId uint) ([]entity.Visit, error)
	SetRemoveVisit(Id uint) (bool, error)
}
type visitConnection struct {
	connection *gorm.DB
}

func NewVisitConnection() VisitRepository {
	return &visitConnection{
		connection: entity.DatabaseConnection(),
	}
}

var errChanVisit = make(chan error, constantvariables.CHAN_VALUE)

/*
@param visit, is a struct of Visit
*/
func (db *visitConnection) SetCreateVisit(visit entity.Visit) (entity.Visit, error) {

	go func() {
		err := db.connection.Save(&visit).Error
		defer entity.Closedb()
		errChanVisit <- err
	}()
	err := <-errChanVisit
	return visit, err
}
func (db *visitConnection) GetAllVisit() ([]entity.Visit, error) {
	var visit []entity.Visit
	go func() {
		err := db.connection.Find(&visit).Error
		errChanVisit <- err
		defer entity.Closedb()
	}()
	err := <-errChanVisit
	return visit, err
}

/*
@param userId, is a uint
@param subDetachmentId, is a uint
*/
func (db *visitConnection) GetAllVisitByUserVisit(userId uint, subDetachmentId uint) ([]entity.Visit, error) {
	var visit []entity.Visit
	go func() {
		err := db.connection.Where("userid=?", userId).
			Where("subdetachmentid=?", subDetachmentId).
			Find(&visit).Error
		errChanVisit <- err
		defer entity.Closedb()
	}()

	err := <-errChanVisit
	return visit, err
}

/*
@param Id, is a uint
*/
func (db *visitConnection) GetFindByIdVisit(Id uint) (entity.Visit, error) {
	var visit entity.Visit
	go func() {
		err := db.connection.Find(&visit, Id).Error
		errChanVisit <- err
		defer entity.Closedb()
	}()

	err := <-errChanVisit
	return visit, err
}

/*
@param Id, is a uint
*/
func (db *visitConnection) SetRemoveVisit(Id uint) (bool, error) {
	go func() {
		err := db.connection.Delete(&entity.Visit{}, Id).Error
		errChanVisit <- err
		defer entity.Closedb()
	}()
	err := <-errChanVisit
	if err != nil {
		return true, err
	}
	return false, err
}
