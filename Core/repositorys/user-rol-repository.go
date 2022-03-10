package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

//UserRepository is contract what UserRepository can do to db
type UserRolRepository interface {
	SetInsertUserRol(role entity.Role) entity.Role
	GetAllUserRole() []entity.Role
}

type userRolConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository

func NewUserRolRepository() UserRolRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &userConnection{
		connection: db,
	}
}

var errChanUserRol = make(chan error, constantvariables.CHAN_VALUE)

/*
@param rol, is a struct of Role
*/
func (db *userConnection) SetInsertUserRol(rol entity.Role) entity.Role {

	go func() {
		err := db.connection.Save(&rol).Error
		defer entity.Closedb()
		errChanUserRol <- err
	}()
	<-errChanUserRol
	return rol
}

func (db *userConnection) GetAllUserRole() []entity.Role {
	var role []entity.Role
	go func() {
		err := db.connection.Joins("User").Joins("Rol").Find(&role).Error
		defer entity.Closedb()
		errChanUserRol <- err
	}()
	<-errChanUserRol
	return role
}
