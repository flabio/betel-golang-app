package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type WeeksDetachmentRepository interface {
	GetFindByIdWeeksDetachment(Id uint) ([]entity.WeeksDetachment, error)
}

type weeksConnection struct {
	connection *gorm.DB
}

func NewWeeksDetachmentRepository() WeeksDetachmentRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &weeksConnection{
		connection: db,
	}
}

/*
@param Id, is a uint
*/
func (db *weeksConnection) GetFindByIdWeeksDetachment(Id uint) ([]entity.WeeksDetachment, error) {
	var weeksdetachment []entity.WeeksDetachment

	err := db.connection.Order("id asc").
		Where("sub_detachment_id", Id).
		Find(&weeksdetachment).Error
	defer entity.Closedb()
	return weeksdetachment, err
}
