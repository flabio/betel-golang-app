package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type ParentRepository interface {
	CreateParent(parent entity.Parent) (entity.Parent, error)
	AllParent() ([]entity.Parent, error)
	AllParentScout(Id uint) ([]entity.Parent, error)
	RemoveParent(Id uint) (bool, error)
	FindParentById(Id uint) (entity.Parent, error)
	FindParentByIdentification(Identification string) (entity.Parent, error)
}

type parentConnection struct {
	connection *gorm.DB
}

func NewParentRepository() ParentRepository {
	return &parentConnection{
		connection: entity.DatabaseConnection(),
	}
}
func (db *parentConnection) CreateParent(parent entity.Parent) (entity.Parent, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&parent).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return parent, err
}
func (db *parentConnection) RemoveParent(Id uint) (bool, error) {
	errChan := make(chan error, 1)

	go func() {
		err := db.connection.Delete(&entity.Parent{}, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err != nil {
		return true, err
	}
	return false, err
}
func (db *parentConnection) FindParentById(Id uint) (entity.Parent, error) {
	errChan := make(chan error, 1)
	var parent entity.Parent
	go func() {
		err := db.connection.Find(&parent, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return parent, err
}
func (db *parentConnection) FindParentByIdentification(Identification string) (entity.Parent, error) {
	errChan := make(chan error, 1)
	var parent entity.Parent
	go func() {
		err := db.connection.Where("identification=?", Identification).Find(&parent).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return parent, err
}

func (db *parentConnection) AllParent() ([]entity.Parent, error) {
	var errChan = make(chan error, 1)
	var parents []entity.Parent
	go func() {
		err := db.connection.Preload("ParentScouts").Find(&parents).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return parents, err
}
func (db *parentConnection) AllParentScout(Id uint) ([]entity.Parent, error) {
	var errChan = make(chan error, 1)
	var parents []entity.Parent
	go func() {
		err := db.connection.Joins("left join parent_scouts on parent_scouts.parent_id = parents.id").Where("parent_scouts.user_id", Id).Find(&parents).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return parents, err
}
