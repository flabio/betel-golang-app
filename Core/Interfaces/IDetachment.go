package Interfaces

import "bete/Core/entity"

type IDetachment interface {
	SetCreateDetachment(detachment entity.Detachment) (entity.Detachment, error)
	SetRemoveDetachment(detachment entity.Detachment) (entity.Detachment, error)
	GetFindDetachmentById(Id uint) (entity.Detachment, error)
	GetAllDetachment() ([]entity.Detachment, error)
}
