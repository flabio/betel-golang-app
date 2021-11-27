package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type PatrolRepository interface {
	Create(patrol entity.Patrol) (entity.Patrol, error)
	Update(patrol entity.Patrol) (entity.Patrol, error)
	Remove(Id uint) (bool, error)
	All() ([]entity.Patrol, error)
	FindById(Id uint) (entity.Patrol, error)
}

type patrolConnection struct {
	connection *gorm.DB
}

func NewPatrolRepository() PatrolRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &patrolConnection{
		connection: db,
	}
}

func (db *patrolConnection) Create(patrol entity.Patrol) (entity.Patrol, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&patrol).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return patrol, err
}

func (db *patrolConnection) Update(patrol entity.Patrol) (entity.Patrol, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&patrol).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return patrol, err
}

func (db *patrolConnection) Remove(Id uint) (bool, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Delete(&entity.Patrol{}, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
func (db *patrolConnection) FindById(Id uint) (entity.Patrol, error) {
	var patrol entity.Patrol
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Find(&patrol, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return patrol, err
}

func (db *patrolConnection) All() ([]entity.Patrol, error) {
	var patrol []entity.Patrol
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Order("id desc").Preload("SubDetachment").Find(&patrol).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return patrol, err
}
