package repositorys

import (
	"bete/Core/entity"

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

/*
@param studycarried, is a struct of StudyCarried
*/
func (db *studycarriedConnection) SetCreateStudyCarried(studycarried entity.StudyCarried) (entity.StudyCarried, error) {

	err := db.connection.Save(&studycarried).Error
	defer entity.Closedb()
	return studycarried, err
}

func (db *studycarriedConnection) GetAllStudyCarried() ([]entity.StudyCarried, error) {
	var result []entity.StudyCarried
	err := db.connection.Find(&result).Error
	defer entity.Closedb()
	return result, err
}

/*
@param Id, is a uint of StudyCarried
*/
func (db *studycarriedConnection) GetFindStudyCarriedById(Id uint) (entity.StudyCarried, error) {
	var result entity.StudyCarried

	err := db.connection.Find(&result, Id).Error
	defer entity.Closedb()
	return result, err
}

/*
@param Id, is a struct of StudyCarried
*/
func (db *studycarriedConnection) SetRemoveStudyCarried(Id uint) (bool, error) {

	err := db.connection.Delete(Id).Error
	defer entity.Closedb()
	if err != nil {
		return true, err
	}
	return false, err
}
