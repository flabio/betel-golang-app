package Interfaces

import "bete/Core/entity"

type IPatrol interface {
	SetCreatePatrol(patrol entity.Patrol) (entity.Patrol, error)
	SetUpdatePatrol(patrol entity.Patrol, Id uint) (entity.Patrol, error)
	SetRemovePatrol(Id uint) (bool, error)
	GetAllPatrol() ([]entity.Patrol, error)
	GetFindByIdPatrol(Id uint) (entity.Patrol, error)
}
