package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

func NewParentRepository() Interfaces.IParent {
	var (
		_OPEN *OpenConnections
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnections{

			connection: entity.Factory(constantvariables.OPTION_FACTORY_DB),
		}
	})
	return _OPEN
}

/*
@param parent, is a struct of Parent
*/
func (db *OpenConnections) SetCreateParent(parent entity.Parent) (entity.Parent, error) {
	db.mux.Lock()
	err := db.connection.Save(&parent).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return parent, err
}

/*
@param Id, is a uint of Parent
*/
func (db *OpenConnections) SetRemoveParent(Id uint) (bool, error) {
	db.mux.Lock()
	err := db.connection.Delete(&entity.Parent{}, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return true, err
	}
	return false, err
}

/*
@param Id, is a uint of Parent
*/
func (db *OpenConnections) GetFindParentById(Id uint) (entity.Parent, error) {
	var parent entity.Parent
	db.mux.Lock()
	err := db.connection.Find(&parent, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return parent, err
}

/*
@param Identification, is a uint of Parent
*/
func (db *OpenConnections) GetFindParentByIdentification(Identification string) (entity.Parent, error) {

	var parent entity.Parent
	db.mux.Lock()
	err := db.connection.Where("identification=?", Identification).Find(&parent).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return parent, err
}

func (db *OpenConnections) GetAllParent() ([]entity.Parent, error) {

	var parents []entity.Parent
	db.mux.Lock()
	err := db.connection.Preload("ParentScouts").Find(&parents).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return parents, err
}

/*
@param Id, is a uint of Parent
*/
func (db *OpenConnections) GetAllParentScout(Id uint) ([]entity.Parent, error) {

	var parents []entity.Parent
	db.mux.Lock()
	err := db.connection.Joins("left join parent_scouts on parent_scouts.parent_id = parents.id").
		Where("parent_scouts.user_id", Id).
		Find(&parents).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return parents, err
}
