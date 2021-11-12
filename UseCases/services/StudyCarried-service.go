package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
)

type StudyCarriedService interface {
	Create(studycarried dto.StudyCarriedDTO) entity.StudyCarried
	AllStudyCarried() []entity.StudyCarried
	FindStudyCarriedById(Id uint) entity.StudyCarried
	DeleteStudyCarried(Id uint) bool
}
type studyCarriedService struct {
	repository repositorys.StudyCarriedRepository
}

//NewStudyCarriedService creates a new instance of StudyCarriedService
func NewStudyCarriedService() StudyCarriedService {
	var repository =repositorys.NewStudyCarriedRepository()
	return &studyCarriedService{
		repository: repository,
	}
}

func (service *studyCarriedService) Create(studyCarriedDto dto.StudyCarriedDTO) entity.StudyCarried {
	studycarriedToCreate := entity.StudyCarried{}
	err := smapping.FillStruct(&studycarriedToCreate, smapping.MapFields(&studyCarriedDto))
	result := service.repository.CreateStudyCarried(studycarriedToCreate)
	checkError(err)
	return result
}

func (service *studyCarriedService) AllStudyCarried() []entity.StudyCarried {
	return service.repository.AllStudyCarried()
}
func (service *studyCarriedService) FindStudyCarriedById(Id uint) entity.StudyCarried {
	return service.repository.FindStudyCarriedById(Id)
}
func (service *studyCarriedService) DeleteStudyCarried(Id uint) bool {
	return service.repository.DeleteStudyCarried(Id)
}
