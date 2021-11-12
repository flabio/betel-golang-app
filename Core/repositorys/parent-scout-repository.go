package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type ParentScoutRepository interface {
	Create(parentscout entity.ParentScout) (entity.ParentScout, error)
	Remove(id int) (bool, error)
}

type parentScoutConnection struct {
	connection *gorm.DB
}

func NewParentScoutRepository() ParentScoutRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &parentScoutConnection{
		connection: db,
	}
}

func (db *parentScoutConnection) Create(parent entity.ParentScout) (entity.ParentScout, error) {

	var errChan = make(chan error, 1)
	go func(db *parentScoutConnection) {
		err := db.connection.Save(&parent).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	err := <-errChan
	return parent, err
}

func (db *parentScoutConnection) Remove(id int) (bool, error) {

	var errChan = make(chan error, 1)
	go func(db *parentScoutConnection) {
		err := db.connection.Delete(&id).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	err := <-errChan
	if err == nil {
		return true, err
	}
	return false, err
}
