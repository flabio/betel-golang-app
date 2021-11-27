package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type StudyCarriedRepository interface {
	CreateStudyCarried(studycarried entity.StudyCarried) entity.StudyCarried
	FindStudyCarriedById(Id uint) entity.StudyCarried
	AllStudyCarried() []entity.StudyCarried
	DeleteStudyCarried(Id uint) bool
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

func (db *studycarriedConnection) CreateStudyCarried(studycarried entity.StudyCarried) entity.StudyCarried {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&studycarried).Error
		defer entity.Closedb()
		errChan <- err
	}()
	<-errChan
	return studycarried
}

func (db *studycarriedConnection) AllStudyCarried() []entity.StudyCarried {
	var result []entity.StudyCarried
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Find(&result).Error
		defer entity.Closedb()
		errChan <- err
	}()
	<-errChan
	return result
}
func (db *studycarriedConnection) FindStudyCarriedById(Id uint) entity.StudyCarried {
	var result entity.StudyCarried
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Find(&result, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	<-errChan
	return result
}
func (db *studycarriedConnection) DeleteStudyCarried(Id uint) bool {

	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Delete(Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	<-errChan
	if <-errChan != nil {
		return true
	}
	return false
}
