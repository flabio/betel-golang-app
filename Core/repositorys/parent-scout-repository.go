package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type ParentScoutRepository interface {
	SetCreateParentScout(parentscout entity.ParentScout) (entity.ParentScout, error)
	SetRemoveParentScout(id int) (bool, error)
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

var errChanParentScout = make(chan error, constantvariables.CHAN_VALUE)

func (db *parentScoutConnection) SetCreateParentScout(parent entity.ParentScout) (entity.ParentScout, error) {

	go func() {
		err := db.connection.Save(&parent).Error
		defer entity.Closedb()
		errChanParentScout <- err
	}()
	err := <-errChanParentScout
	return parent, err
}

func (db *parentScoutConnection) SetRemoveParentScout(id int) (bool, error) {

	go func() {
		err := db.connection.Delete(&id).Error
		defer entity.Closedb()
		errChanParentScout <- err
	}()
	err := <-errChanParentScout
	if err == nil {
		return true, err
	}
	return false, err
}
