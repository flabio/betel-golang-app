package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type CityRepository interface {
	GetAllCity() ([]entity.City, error)
}

type cityConnection struct {
	connection *gorm.DB
}

func NewCityRepository() CityRepository {
	return &cityConnection{
		connection: entity.DatabaseConnection(),
	}
}

func (db *cityConnection) GetAllCity() ([]entity.City, error) {

	var citys []entity.City
	err := db.connection.Find(&citys).Error
	defer entity.Closedb()
	return citys, err
}
