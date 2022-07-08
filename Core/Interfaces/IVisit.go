package Interfaces

import "bete/Core/entity"

type IVisit interface {
	SetCreateVisit(visit entity.Visit) (entity.Visit, error)
	GetAllVisit() ([]entity.Visit, error)
	GetFindByIdVisit(Id uint) (entity.Visit, error)
	GetAllVisitByUserVisit(userId uint, subDetachmentId uint) ([]entity.Visit, error)
	SetRemoveVisit(Id uint) (bool, error)
}
