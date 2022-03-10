package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

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

var errChanWeeksDetachment = make(chan error, constantvariables.CHAN_VALUE)

/*
@param Id, is a uint
*/
func (db *weeksConnection) GetFindByIdWeeksDetachment(Id uint) ([]entity.WeeksDetachment, error) {
	var weeksdetachment []entity.WeeksDetachment

	go func() {
		err := db.connection.Order("id asc").
			Where("sub_detachment_id", Id).
			Find(&weeksdetachment).Error
		defer entity.Closedb()
		errChanWeeksDetachment <- err
	}()
	err := <-errChanWeeksDetachment
	return weeksdetachment, err
}
