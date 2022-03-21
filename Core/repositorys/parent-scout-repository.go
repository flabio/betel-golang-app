package repositorys

import (
	"bete/Core/entity"

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

/*
@param parent, is a struct of ParentScout
*/
func (db *parentScoutConnection) SetCreateParentScout(parent entity.ParentScout) (entity.ParentScout, error) {

	err := db.connection.Save(&parent).Error
	defer entity.Closedb()

	return parent, err
}

/*
@param Id, is a uint of ParentScout
*/
func (db *parentScoutConnection) SetRemoveParentScout(id int) (bool, error) {

	err := db.connection.Delete(&id).Error
	defer entity.Closedb()

	if err == nil {
		return true, err
	}
	return false, err
}
