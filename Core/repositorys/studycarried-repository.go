package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type StudyCarriedRepository interface {
	SetCreateStudyCarried(studycarried entity.StudyCarried) (entity.StudyCarried, error)
	GetFindStudyCarriedById(Id uint) (entity.StudyCarried, error)
	GetAllStudyCarried() ([]entity.StudyCarried, error)
	SetRemoveStudyCarried(Id uint) (bool, error)
}

type studycarriedConnection struct {
	connection *gorm.DB
}

func NewStudyCarriedRepository() StudyCarriedRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &studycarriedConnection{
		connection: db,
	}
}

var errChanStudyCarried = make(chan error, constantvariables.CHAN_VALUE)

func (db *studycarriedConnection) SetCreateStudyCarried(studycarried entity.StudyCarried) (entity.StudyCarried, error) {

	go func() {
		err := db.connection.Save(&studycarried).Error
		defer entity.Closedb()
		errChanStudyCarried <- err
	}()
	err := <-errChanStudyCarried
	return studycarried, err
}

func (db *studycarriedConnection) GetAllStudyCarried() ([]entity.StudyCarried, error) {
	var result []entity.StudyCarried
	go func() {
		err := db.connection.Find(&result).Error
		defer entity.Closedb()
		errChanStudyCarried <- err
	}()
	err := <-errChanStudyCarried
	return result, err
}
func (db *studycarriedConnection) GetFindStudyCarriedById(Id uint) (entity.StudyCarried, error) {
	var result entity.StudyCarried

	go func() {
		err := db.connection.Find(&result, Id).Error
		defer entity.Closedb()
		errChanStudyCarried <- err
	}()
	err := <-errChanStudyCarried
	return result, err
}
func (db *studycarriedConnection) SetRemoveStudyCarried(Id uint) (bool, error) {

	go func() {
		err := db.connection.Delete(Id).Error
		defer entity.Closedb()
		errChanStudyCarried <- err
	}()
	err := <-errChanStudyCarried
	if <-errChanStudyCarried != nil {
		return true, err
	}
	return false, err
}
