package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type ScoutParentRepository interface {
	CreateParentScout(parentScout entity.ParentScout) (entity.ParentScout, error)
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

func (db *parentscoutConnection) CreateParentScout(parentScout entity.ParentScout) (entity.ParentScout, error) {
	err := db.connection.Save(&parentScout).Error

	return parentScout, err
}
