package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type SubDetachmentRepository interface {
	SetCreateSubDetachment(subdetachment entity.SubDetachment) (entity.SubDetachment, error)
	SetRemoveSubDetachment(Id uint) (bool, error)
	GetFindByIdSubDetachment(Id uint) (entity.SubDetachment, error)
	GetFindByIdDetachmentSubDetachment(Id uint) ([]entity.SubDetachment, error)
	GetAllSubDetachment() ([]entity.SubDetachment, error)
}

type subConnection struct {
	connection *gorm.DB
}

func NewSubDetachmentRepository() SubDetachmentRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &subConnection{
		connection: db,
	}
}

/*
@param subdetachment, is a struct of SubDetachment
*/
func (db *subConnection) SetCreateSubDetachment(subdetachment entity.SubDetachment) (entity.SubDetachment, error) {
	err := db.connection.Save(&subdetachment).Error
	defer entity.Closedb()
	return subdetachment, err
}

/*
@param Id, is a uint of SubDetachment
*/
func (db *subConnection) SetRemoveSubDetachment(Id uint) (bool, error) {
	err := db.connection.Delete(&entity.SubDetachment{}, Id).Error
	defer entity.Closedb()
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

/*
@param Id, is a uint of SubDetachment
*/
func (db *subConnection) GetFindByIdSubDetachment(Id uint) (entity.SubDetachment, error) {
	var subdetachment entity.SubDetachment
	err := db.connection.Find(&subdetachment, Id).Error
	defer entity.Closedb()
	return subdetachment, err
}
func (db *subConnection) GetAllSubDetachment() ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	err := db.connection.Order("id desc").Preload("Detachment").Find(&subdetachment).Error
	defer entity.Closedb()
	return subdetachment, err
}

/*
@param Id, is a uint of SubDetachment
*/
func (db *subConnection) GetFindByIdDetachmentSubDetachment(Id uint) ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	err := db.connection.Order("id desc").Preload("Detachment").Where("detachment_id", Id).Find(&subdetachment).Error
	defer entity.Closedb()
	return subdetachment, err
}
