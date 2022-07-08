package Interfaces

import "bete/Core/entity"

type ISubDetachment interface {
	SetCreateSubDetachment(subdetachment entity.SubDetachment) (entity.SubDetachment, error)
	SetRemoveSubDetachment(Id uint) (bool, error)
	GetFindByIdSubDetachment(Id uint) (entity.SubDetachment, error)
	GetFindByIdDetachmentSubDetachment(Id uint) ([]entity.SubDetachment, error)
	GetAllSubDetachment() ([]entity.SubDetachment, error)
}
