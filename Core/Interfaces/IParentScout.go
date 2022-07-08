package Interfaces

import "bete/Core/entity"

type IParentScout interface {
	SetCreateParentScout(parentscout entity.ParentScout) (entity.ParentScout, error)
	SetRemoveParentScout(id int) (bool, error)
}
