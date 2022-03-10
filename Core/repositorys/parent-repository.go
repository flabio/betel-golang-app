package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type ParentRepository interface {
	SetCreateParent(parent entity.Parent) (entity.Parent, error)
	GetAllParent() ([]entity.Parent, error)
	GetAllParentScout(Id uint) ([]entity.Parent, error)
	SetRemoveParent(Id uint) (bool, error)
	GetFindParentById(Id uint) (entity.Parent, error)
	GetFindParentByIdentification(Identification string) (entity.Parent, error)
}

type parentConnection struct {
	connection *gorm.DB
}

func NewParentRepository() ParentRepository {
	return &parentConnection{
		connection: entity.DatabaseConnection(),
	}
}

var errChanParent = make(chan error, constantvariables.CHAN_VALUE)

/*
@param parent, is a struct of Parent
*/
func (db *parentConnection) SetCreateParent(parent entity.Parent) (entity.Parent, error) {

	go func() {
		err := db.connection.Save(&parent).Error
		defer entity.Closedb()
		errChanParent <- err
	}()
	err := <-errChanParent
	return parent, err
}

/*
@param Id, is a uint of Parent
*/
func (db *parentConnection) SetRemoveParent(Id uint) (bool, error) {
	go func() {
		err := db.connection.Delete(&entity.Parent{}, Id).Error
		defer entity.Closedb()
		errChanParent <- err
	}()
	err := <-errChanParent
	if err != nil {
		return true, err
	}
	return false, err
}

/*
@param Id, is a uint of Parent
*/
func (db *parentConnection) GetFindParentById(Id uint) (entity.Parent, error) {
	var parent entity.Parent
	go func() {
		err := db.connection.Find(&parent, Id).Error
		defer entity.Closedb()
		errChanParent <- err
	}()
	err := <-errChanParent
	return parent, err
}

/*
@param Identification, is a uint of Parent
*/
func (db *parentConnection) GetFindParentByIdentification(Identification string) (entity.Parent, error) {

	var parent entity.Parent
	go func() {
		err := db.connection.Where("identification=?", Identification).Find(&parent).Error
		defer entity.Closedb()
		errChanParent <- err
	}()
	err := <-errChanParent
	return parent, err
}

func (db *parentConnection) GetAllParent() ([]entity.Parent, error) {

	var parents []entity.Parent
	go func() {
		err := db.connection.Preload("ParentScouts").Find(&parents).Error
		defer entity.Closedb()
		errChanParent <- err
	}()
	err := <-errChanParent
	return parents, err
}

/*
@param Id, is a uint of Parent
*/
func (db *parentConnection) GetAllParentScout(Id uint) ([]entity.Parent, error) {

	var parents []entity.Parent
	go func() {
		err := db.connection.Joins("left join parent_scouts on parent_scouts.parent_id = parents.id").
			Where("parent_scouts.user_id", Id).
			Find(&parents).Error
		defer entity.Closedb()
		errChanParent <- err
	}()
	err := <-errChanParent
	return parents, err
}
