package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	SetCreateAttendance(attendance entity.Attendance) (entity.Attendance, error)
	SetRemoveAttendance(Id uint) (bool, error)
	GetFindByIdAttendance(Id uint) (entity.Attendance, error)
	GetFindByIdWeeksDetachment(Week string, IdUser uint) (entity.Attendance, error)
	GetAllAttendance() ([]entity.Attendance, error)
	GetAttendancesSubdetachment(IdUser uint, IdSubDetachment uint) ([]entity.Attendance, error)
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

//NewAttendanceRepository()

func (db *attendanceConnection) SetCreateAttendance(attendance entity.Attendance) (entity.Attendance, error) {
	err := db.connection.Save(&attendance).Error
	defer entity.Closedb()
	return attendance, err
}

//SetCreateAttendance

/*
@param Id is Attendace ,of type uint
*/
func (db *attendanceConnection) SetRemoveAttendance(Id uint) (bool, error) {
	var attendance entity.Attendance
	err := db.connection.Delete(&attendance, Id).Error
	defer entity.Closedb()

	if err != nil {
		return false, err
	}
	return true, err
}
func (db *attendanceConnection) GetAllAttendance() ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	err := db.connection.Preload("User").Preload("SubDetachment").Find(&attendance).Error
	defer entity.Closedb()
	return attendance, err
}

/*
@param IdUser is the user, of type uint
@param IdSubDetachment is the Subdetachment, of type uint
*/
func (db *attendanceConnection) GetAttendancesSubdetachment(IdUser uint, IdSubDetachment uint) ([]entity.Attendance, error) {
	var attendance []entity.Attendance

	err := db.connection.Preload("User").
		Preload("SubDetachment").
		Where("user_id=?", IdUser).
		Where("sub_detachment_id=?", IdSubDetachment).
		Find(&attendance).Error
	defer entity.Closedb()
	return attendance, err
}

/*
@param Id is the attendance, of type uint
*/
func (db *attendanceConnection) GetFindByIdAttendance(Id uint) (entity.Attendance, error) {
	var attendance entity.Attendance
	err := db.connection.Find(&attendance, Id).Error
	defer entity.Closedb()
	return attendance, err
}

/*
@param Week is the attendance, of type uint
@param IdUser is the attendance, of type uint

*/
func (db *attendanceConnection) GetFindByIdWeeksDetachment(Week string, IdUser uint) (entity.Attendance, error) {
	var attendance entity.Attendance

	err := db.connection.Where("user_id=?", IdUser).
		Where("week_sub_detachment=?", Week).
		Find(&attendance).Error
	defer entity.Closedb()

	return attendance, err
}
