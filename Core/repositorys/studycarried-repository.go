package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func NewStudyCarriedRepository() Interfaces.IStudyCarried {
	var (
		_OPEN *OpenConnections
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnections{

			connection: entity.Factory(constantvariables.OPTION_FACTORY_DB),
		}
	})
	return _OPEN
}

/*
@param studycarried, is a struct of StudyCarried
*/
func (db *OpenConnections) SetCreateStudyCarried(studycarried entity.StudyCarried) (entity.StudyCarried, error) {
	db.mux.Lock()
	err := db.connection.Save(&studycarried).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return studycarried, err
}

func (db *OpenConnections) GetAllStudyCarried() ([]entity.StudyCarried, error) {
	var result []entity.StudyCarried
	db.mux.Lock()
	err := db.connection.Find(&result).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return result, err
}

/*
@param Id, is a uint of StudyCarried
*/
func (db *OpenConnections) GetFindStudyCarriedById(Id uint) (entity.StudyCarried, error) {
	var result entity.StudyCarried
	db.mux.Lock()
	err := db.connection.Find(&result, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return result, err
}

/*
@param Id, is a struct of StudyCarried
*/
func (db *OpenConnections) SetRemoveStudyCarried(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return true, err
	}
	return false, err
}
