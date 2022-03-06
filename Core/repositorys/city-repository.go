package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

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

var errChanCity = make(chan error, constantvariables.CHAN_VALUE)

func (db *cityConnection) GetAllCity() ([]entity.City, error) {

	var citys []entity.City
	go func() {
		err := db.connection.Find(&citys).Error
		defer entity.Closedb()
		errChanCity <- err
	}()
	err := <-errChanCity
	return citys, err
}
