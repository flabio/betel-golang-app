package repositorys

import (
	"bete/Core/entity"

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

/*
@param parentScout, is a struct of ParentScout
*/
func (db *parentscoutConnection) SetCreateParentScout(parentScout entity.ParentScout) (entity.ParentScout, error) {
	err := db.connection.Save(&parentScout).Error
	defer entity.Closedb()
	return parentScout, err
}
