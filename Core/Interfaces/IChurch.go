package Interfaces

import "bete/Core/entity"

type IChurch interface {
	SetCreateChurch(church entity.Church) (entity.Church, error)
	SetRemoveChurch(church entity.Church) (bool, error)
	GetFindChurchById(Id uint) (entity.Church, error)
	GetAllChurch() ([]entity.Church, error)
}
