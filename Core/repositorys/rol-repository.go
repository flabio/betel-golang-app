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

/*
@param rol, is a struct of Rol
*/
func (db *rolConnection) SetCreateRol(rol entity.Rol) (entity.Rol, error) {

	err := db.connection.Save(&rol).Error
	defer entity.Closedb()
	return rol, err
}

/*
@param rol, is a struct of Rol
*/
func (db *rolConnection) SetRemoveRol(rol entity.Rol) (bool, error) {

	err := db.connection.Delete(&rol).Error
	defer entity.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}

/*
@param Id, is a uint of Rol
*/
func (db *rolConnection) GetFindRolById(Id uint) (entity.Rol, error) {

	var rol entity.Rol
	err := db.connection.Find(&rol, Id).Error
	defer entity.Closedb()
	return rol, err
}

func (db *rolConnection) GetAllRol() ([]entity.Rol, error) {

	var rols []entity.Rol
	err := db.connection.Find(&rols).Error
	defer entity.Closedb()
	return rols, err
}

func (db *rolConnection) GetAllGroupRol() ([]entity.Rol, error) {
	var rols []entity.Rol

	err := db.connection.Where("id IN ?", []int{
		constantvariables.NAVIGANTORS_ROL,
		constantvariables.PIONEERS_ROL,
		constantvariables.PATH_FOLLOWERS_ROL,
		constantvariables.SCOUTS_ROL}).
		Find(&rols).Error
	defer entity.Closedb()
	return rols, err
}
func (db *rolConnection) GetRolsModule() ([]entity.RoleModule, error) {

	var roleModule []entity.RoleModule

	err := db.connection.Preload("Role.Rol").
		Preload("Module").
		Find(&roleModule).Error
	defer entity.Closedb()

	return roleModule, err
}
