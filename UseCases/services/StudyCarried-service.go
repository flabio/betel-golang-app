package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
)

type studyCarriedService struct {
	IStudyCarried Interfaces.IStudyCarried
}

//NewStudyCarriedService creates a new instance of StudyCarriedService
func NewStudyCarriedService() InterfacesService.IStudyCarriedService {
	return &studyCarriedService{
		IStudyCarried: repositorys.NewStudyCarriedRepository(),
	}
}

func (studyCarried *studyCarriedService) SetCreateStudyCarriedService(studyCarriedDto dto.StudyCarriedDTO) entity.StudyCarried {
	studycarriedToCreate := entity.StudyCarried{}
	err := smapping.FillStruct(&studycarriedToCreate, smapping.MapFields(&studyCarriedDto))
	result, err := studyCarried.IStudyCarried.SetCreateStudyCarried(studycarriedToCreate)
	checkError(err)
	return result
}

func (studyCarried *studyCarriedService) GetAllStudyCarriedService() []entity.StudyCarried {
	result, _ := studyCarried.IStudyCarried.GetAllStudyCarried()
	return result
}
func (studyCarried *studyCarriedService) GetFindStudyCarriedByIdService(Id uint) entity.StudyCarried {
	result, _ := studyCarried.IStudyCarried.GetFindStudyCarriedById(Id)
	return result
}
func (studyCarried *studyCarriedService) SetRemoveStudyCarriedService(Id uint) bool {
	result, _ := studyCarried.IStudyCarried.SetRemoveStudyCarried(Id)
	return result
}
