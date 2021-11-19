package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type SubDetachmentRepository interface {
	Create(subdetachment entity.SubDetachment) (entity.SubDetachment, error)
	Update(subdetachment entity.SubDetachment) (entity.SubDetachment, error)
	Remove(Id uint) (bool, error)
	FindById(Id uint) (entity.SubDetachment, error)
	FindByIdDetachment(Id uint) ([]entity.SubDetachment, error)
	All() ([]entity.SubDetachment, error)

	AddUserSubDetachment(userSubDetachment entity.UserSubdetachement) (entity.UserSubdetachement, error)
	RemoveUserSubDetachment(Id uint) (bool, error)
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

func (db *subConnection) Create(subdetachment entity.SubDetachment) (entity.SubDetachment, error) {
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Save(&subdetachment).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	return subdetachment, <-errChan
}

func (db *subConnection) Update(subdetachment entity.SubDetachment) (entity.SubDetachment, error) {
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Save(&subdetachment).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	return subdetachment, <-errChan
}

func (db *subConnection) Remove(Id uint) (bool, error) {
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Delete(&entity.SubDetachment{}, Id).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	err := <-errChan
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
func (db *subConnection) FindById(Id uint) (entity.SubDetachment, error) {
	var subdetachment entity.SubDetachment
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Find(&subdetachment, Id).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	err := <-errChan
	return subdetachment, err
}
func (db *subConnection) All() ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Order("id desc").Preload("Detachment").Find(&subdetachment).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	err := <-errChan
	return subdetachment, err
}
func (db *subConnection) FindByIdDetachment(Id uint) ([]entity.SubDetachment, error) {
	var subdetachment []entity.SubDetachment
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Order("id desc").Preload("Detachment").Where("detachment_id", Id).Find(&subdetachment).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	err := <-errChan
	return subdetachment, err
}

//AddUserSubDetachment
func (db *subConnection) AddUserSubDetachment(userSubdetachement entity.UserSubdetachement) (entity.UserSubdetachement, error) {
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Save(&userSubdetachement).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	return userSubdetachement, <-errChan
}

func (db *subConnection) RemoveUserSubDetachment(Id uint) (bool, error) {
	var errChan = make(chan error, 1)
	go func(db *subConnection) {
		err := db.connection.Delete(&entity.UserSubdetachement{}, Id).Error
		defer entity.Closedb()
		errChan <- err
	}(db)
	err := <-errChan
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
