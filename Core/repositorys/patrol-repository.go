package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type PatrolRepository interface {
	SetCreatePatrol(patrol entity.Patrol) (entity.Patrol, error)
	SetRemovePatrol(Id uint) (bool, error)
	GetAllPatrol() ([]entity.Patrol, error)
	GetFindByIdPatrol(Id uint) (entity.Patrol, error)
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

var errChanPatrol = make(chan error, constantvariables.CHAN_VALUE)

/*
@param patrol, is a struct of Patrol
*/
func (db *patrolConnection) SetCreatePatrol(patrol entity.Patrol) (entity.Patrol, error) {
	go func() {
		err := db.connection.Save(&patrol).Error
		defer entity.Closedb()
		errChanPatrol <- err
	}()
	err := <-errChanPatrol
	return patrol, err
}

/*
@param Id, is a uint of Patrol
*/
func (db *patrolConnection) SetRemovePatrol(Id uint) (bool, error) {
	go func() {
		err := db.connection.Delete(&entity.Patrol{}, Id).Error
		defer entity.Closedb()
		errChanPatrol <- err
	}()
	err := <-errChanPatrol
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

/*
@param Id, is a uint of Patrol
*/
func (db *patrolConnection) GetFindByIdPatrol(Id uint) (entity.Patrol, error) {
	var patrol entity.Patrol
	go func() {
		err := db.connection.Find(&patrol, Id).Error
		defer entity.Closedb()
		errChanPatrol <- err
	}()
	err := <-errChanPatrol
	return patrol, err
}

func (db *patrolConnection) GetAllPatrol() ([]entity.Patrol, error) {
	var patrol []entity.Patrol
	go func() {
		err := db.connection.Order("id desc").Preload("SubDetachment").Find(&patrol).Error
		defer entity.Closedb()
		errChanPatrol <- err
	}()
	err := <-errChanPatrol
	return patrol, err
}
