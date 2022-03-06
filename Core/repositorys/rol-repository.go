package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type RolRepository interface {
	SetCreateRol(rol entity.Rol) (entity.Rol, error)
	GetAllRol() ([]entity.Rol, error)
	GetAllGroupRol() ([]entity.Rol, error)
	GetRolsModule() ([]entity.RoleModule, error)
	SetRemoveRol(rol entity.Rol) (bool, error)
	GetFindRolById(Id uint) (entity.Rol, error)
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

var errChanRol = make(chan error, constantvariables.CHAN_VALUE)

func (db *rolConnection) SetCreateRol(rol entity.Rol) (entity.Rol, error) {

	go func() {
		err := db.connection.Save(&rol).Error
		defer entity.Closedb()
		errChanRol <- err
	}()
	err := <-errChanRol

	return rol, err
}
func (db *rolConnection) SetRemoveRol(rol entity.Rol) (bool, error) {

	go func() {
		err := db.connection.Delete(&rol).Error
		defer entity.Closedb()
		errChanRol <- err
	}()
	err := <-errChanRol
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *rolConnection) GetFindRolById(Id uint) (entity.Rol, error) {

	var rol entity.Rol
	go func() {
		err := db.connection.Find(&rol, Id).Error
		defer entity.Closedb()
		errChanRol <- err
	}()
	err := <-errChanRol
	return rol, err
}

func (db *rolConnection) GetAllRol() ([]entity.Rol, error) {

	var rols []entity.Rol
	go func() {
		err := db.connection.Find(&rols).Error
		defer entity.Closedb()
		errChanRol <- err
	}()
	err := <-errChanRol
	return rols, err
}

func (db *rolConnection) GetAllGroupRol() ([]entity.Rol, error) {
	var rols []entity.Rol

	go func() {
		err := db.connection.Where("id IN ?", []int{
			constantvariables.NAVIGANTORS_ROL,
			constantvariables.PIONEERS_ROL,
			constantvariables.PATH_FOLLOWERS_ROL,
			constantvariables.SCOUTS_ROL}).
			Find(&rols).Error
		defer entity.Closedb()
		errChanRol <- err

	}()
	err := <-errChanRol
	return rols, err
}
func (db *rolConnection) GetRolsModule() ([]entity.RoleModule, error) {

	var roleModule []entity.RoleModule

	go func() {
		err := db.connection.Preload("Role.Rol").
			Preload("Module").
			Find(&roleModule).Error
		defer entity.Closedb()
		errChanRol <- err
	}()
	err := <-errChanRol
	return roleModule, err
}
