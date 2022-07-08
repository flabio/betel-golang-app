package Interfaces

import "bete/Core/entity"

type IScoutParent interface {
	SetCreateParentScouts(parentScout entity.ParentScout) (entity.ParentScout, error)
}
