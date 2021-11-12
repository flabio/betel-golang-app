package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

//UserRepository is contract what UserRepository can do to db
type UserRolRepository interface {
	InsertUserRol(role entity.Role) entity.Role
	AllUserRole() []entity.Role
}

type userRolConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository

func NewUserRolRepository() UserRolRepository {
	var db *gorm.DB=entity.DatabaseConnection()
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUserRol(rol entity.Role) entity.Role {
	var errChan=make(chan error,1)
	go func(db *userConnection){
		err:=db.connection.Save(&rol).Error
		defer entity.Closedb()
		errChan<-err
	}(db)
	<-errChan
	return rol
}

func (db *userConnection) AllUserRole() []entity.Role {
	var role []entity.Role
	var errChan=make(chan error,1)
	go func(db *userConnection){
		err:=db.connection.Joins("User").Joins("Rol").Find(&role).Error
		defer entity.Closedb()
		errChan<-err
	}(db)
	<-errChan
	return role
}
