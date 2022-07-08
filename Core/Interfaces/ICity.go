package Interfaces

import "bete/Core/entity"

type ICity interface {
	GetAllCity() ([]entity.City, error)
}
