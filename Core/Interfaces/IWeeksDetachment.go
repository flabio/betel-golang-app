package Interfaces

import "bete/Core/entity"

type IWeeksDetachment interface {
	GetFindByIdWeeks(Id uint) ([]entity.WeeksDetachment, error)
}
