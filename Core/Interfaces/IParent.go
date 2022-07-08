package Interfaces

import "bete/Core/entity"

type IParent interface {
	SetCreateParent(parent entity.Parent) (entity.Parent, error)
	GetAllParent() ([]entity.Parent, error)
	GetAllParentScout(Id uint) ([]entity.Parent, error)
	SetRemoveParent(Id uint) (bool, error)
	GetFindParentById(Id uint) (entity.Parent, error)
	GetFindParentByIdentification(Identification string) (entity.Parent, error)
}
