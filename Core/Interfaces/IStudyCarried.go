package Interfaces

import "bete/Core/entity"

type IStudyCarried interface {
	SetCreateStudyCarried(studycarried entity.StudyCarried) (entity.StudyCarried, error)
	GetFindStudyCarriedById(Id uint) (entity.StudyCarried, error)
	GetAllStudyCarried() ([]entity.StudyCarried, error)
	SetRemoveStudyCarried(Id uint) (bool, error)
}
