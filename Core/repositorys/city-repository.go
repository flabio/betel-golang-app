package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type CityRepository interface {
	AllCity() ([]entity.City, error)
}

type cityConnection struct {
	connection *gorm.DB
}

func NewCityRepository() CityRepository {
	return &cityConnection{
		connection: entity.DatabaseConnection(),
	}
}
func (db *cityConnection) AllCity() ([]entity.City, error) {
	var errChan = make(chan error, 1)
	var citys []entity.City
	go func() {
		err := db.connection.Find(&citys).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return citys, err
}
