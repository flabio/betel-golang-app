package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type ScoutParentRepository interface {
	SetCreateParentScout(parentScout entity.ParentScout) (entity.ParentScout, error)
}

type parentscoutConnection struct {
	connection *gorm.DB
}

func NewScoutParentRepository() ScoutParentRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &parentscoutConnection{
		connection: db,
	}
}

var errChanParentS = make(chan error, constantvariables.CHAN_VALUE)

func (db *parentscoutConnection) SetCreateParentScout(parentScout entity.ParentScout) (entity.ParentScout, error) {
	go func() {
		err := db.connection.Save(&parentScout).Error
		defer entity.Closedb()
		errChanParentS <- err
	}()
	err := <-errChanParentS
	return parentScout, err
}
