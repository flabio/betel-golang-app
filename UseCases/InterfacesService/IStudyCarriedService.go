package InterfacesService

import (
	"bete/Core/entity"
	"bete/UseCases/dto"
)

type IStudyCarriedService interface {
	SetCreateStudyCarriedService(studycarried dto.StudyCarriedDTO) entity.StudyCarried
	GetAllStudyCarriedService() []entity.StudyCarried
	GetFindStudyCarriedByIdService(Id uint) entity.StudyCarried
	SetRemoveStudyCarriedService(Id uint) bool
}
