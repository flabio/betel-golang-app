package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
)

type StudyCarriedService interface {
	SetCreateStudyCarriedService(studycarried dto.StudyCarriedDTO) entity.StudyCarried
	GetAllStudyCarriedService() []entity.StudyCarried
	GetFindStudyCarriedByIdService(Id uint) entity.StudyCarried
	SetRemoveStudyCarriedService(Id uint) bool
}
type studyCarriedService struct {
	repository repositorys.StudyCarriedRepository
}

//NewStudyCarriedService creates a new instance of StudyCarriedService
func NewStudyCarriedService() StudyCarriedService {
	var repository = repositorys.NewStudyCarriedRepository()
	return &studyCarriedService{
		repository: repository,
	}
}

func (service *studyCarriedService) SetCreateStudyCarriedService(studyCarriedDto dto.StudyCarriedDTO) entity.StudyCarried {
	studycarriedToCreate := entity.StudyCarried{}
	err := smapping.FillStruct(&studycarriedToCreate, smapping.MapFields(&studyCarriedDto))
	result, err := service.repository.SetCreateStudyCarried(studycarriedToCreate)
	checkError(err)
	return result
}

func (service *studyCarriedService) GetAllStudyCarriedService() []entity.StudyCarried {
	result, _ := service.repository.GetAllStudyCarried()
	return result
}
func (service *studyCarriedService) GetFindStudyCarriedByIdService(Id uint) entity.StudyCarried {
	result, _ := service.repository.GetFindStudyCarriedById(Id)
	return result
}
func (service *studyCarriedService) SetRemoveStudyCarriedService(Id uint) bool {
	result, _ := service.repository.SetRemoveStudyCarried(Id)
	return result
}
