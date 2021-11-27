package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(attendance entity.Attendance) (entity.Attendance, error)
	Update(attendance entity.Attendance) (entity.Attendance, error)
	Remove(Id uint) (bool, error)
	FindByIdAttendance(Id uint) (entity.Attendance, error)
	All() ([]entity.Attendance, error)
}

type attendanceConnection struct {
	connection *gorm.DB
}

func NewAttendanceRepository() AttendanceRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &attendanceConnection{
		connection: db,
	}

}
func (db *attendanceConnection) Create(attendance entity.Attendance) (entity.Attendance, error) {
	err := db.connection.Save(&attendance).Error
	defer entity.Closedb()
	return attendance, err
}
func (db *attendanceConnection) Update(attendance entity.Attendance) (entity.Attendance, error) {
	err := db.connection.Save(&attendance).Error
	defer entity.Closedb()
	return attendance, err
}
func (db *attendanceConnection) Remove(Id uint) (bool, error) {
	var attendance entity.Attendance
	err := db.connection.Delete(&attendance, Id).Error
	defer entity.Closedb()
	if err != nil {
		return false, err
	}
	return true, err
}
func (db *attendanceConnection) All() ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	err := db.connection.Find(&attendance).Error
	defer entity.Closedb()
	return attendance, err
}
func (db *attendanceConnection) FindByIdAttendance(Id uint) (entity.Attendance, error) {
	var attendance entity.Attendance
	err := db.connection.Find(&attendance, Id).Error
	defer entity.Closedb()
	return attendance, err
}