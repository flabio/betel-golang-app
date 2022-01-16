package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type WeeksDetachmentRepository interface {
	FindByIdWeeksDetachment(Id uint) ([]entity.WeeksDetachment, error)
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

func (db *weeksConnection) FindByIdWeeksDetachment(Id uint) ([]entity.WeeksDetachment, error) {
	var weeksdetachment []entity.WeeksDetachment
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Order("id asc").Where("sub_detachment_id", Id).Find(&weeksdetachment).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return weeksdetachment, err
}
