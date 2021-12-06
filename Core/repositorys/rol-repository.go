package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type RolRepository interface {
	CreateRol(rol entity.Rol) (entity.Rol, error)
	UpdateRol(rol entity.Rol) (entity.Rol, error)
	AllRol() ([]entity.Rol, error)
	AllGroupRol() ([]entity.Rol, error)
	RolsModule() ([]entity.RoleModule, error)
	DeleteRol(rol entity.Rol) (bool, error)
	FindRolById(Id uint) (entity.Rol, error)
}
type rolConnection struct {
	connection *gorm.DB
}

func NewRolRepository() RolRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &rolConnection{
		connection: db,
	}
}

func (db *rolConnection) CreateRol(rol entity.Rol) (entity.Rol, error) {
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Save(&rol).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan

	return rol, err
}
func (db *rolConnection) UpdateRol(rol entity.Rol) (entity.Rol, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&rol).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return rol, err
}
func (db *rolConnection) DeleteRol(rol entity.Rol) (bool, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Delete(&rol).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *rolConnection) FindRolById(Id uint) (entity.Rol, error) {

	var rol entity.Rol
	var errChan = make(chan error)
	go func() {
		err := db.connection.Find(&rol, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return rol, err
}

func (db *rolConnection) AllRol() ([]entity.Rol, error) {
	var errChan = make(chan error, 1)
	var rols []entity.Rol
	go func() {
		err := db.connection.Find(&rols).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return rols, err
}

func (db *rolConnection) AllGroupRol() ([]entity.Rol, error) {
	var rols []entity.Rol
	err := db.connection.Where("id IN ?", []int{6, 26, 27, 28}).Find(&rols).Error
	defer entity.Closedb()
	return rols, err
}
func (db *rolConnection) RolsModule() ([]entity.RoleModule, error) {
	var errChan = make(chan error, 1)
	var roleModule []entity.RoleModule

	go func() {
		err := db.connection.Preload("Role.Rol").Preload("Module").Find(&roleModule).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return roleModule, err
}
