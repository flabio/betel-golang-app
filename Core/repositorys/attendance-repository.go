package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	SetCreateAttendance(attendance entity.Attendance) (entity.Attendance, error)
	SetRemoveAttendance(Id uint) (bool, error)
	GetFindByIdAttendance(Id uint) (entity.Attendance, error)
	GetFindByIdWeeksDetachment(week string, userid uint) (entity.Attendance, error)
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

var errChan = make(chan error, constantvariables.CHAN_VALUE)

func (db *attendanceConnection) SetCreateAttendance(attendance entity.Attendance) (entity.Attendance, error) {

	go func() {
		err := db.connection.Save(&attendance).Error
		defer entity.Closedb()
		errChan <- err
	}()
	//go
	err := <-errChan

	return attendance, err
}

//SetCreateAttendance

/*
@param Id is Attendace ,of type uint
*/
func (db *attendanceConnection) SetRemoveAttendance(Id uint) (bool, error) {
	var attendance entity.Attendance
	go func() {
		err := db.connection.Delete(&attendance, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err != nil {
		return false, err
	}
	return true, err
}
func (db *attendanceConnection) GetAllAttendance() ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	go func() {
		err := db.connection.Preload("User").Preload("SubDetachment").Find(&attendance).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return attendance, err
}
func (db *attendanceConnection) GetAttendancesSubdetachment(IdUser uint, IdSubDetachment uint) ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	go func() {
		err := db.connection.Preload("User").
			Preload("SubDetachment").
			Where("user_id=?", IdUser).
			Where("sub_detachment_id=?", IdSubDetachment).
			Find(&attendance).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return attendance, err
}
func (db *attendanceConnection) GetFindByIdAttendance(Id uint) (entity.Attendance, error) {
	var attendance entity.Attendance
	go func() {
		err := db.connection.Find(&attendance, Id).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return attendance, err
}

func (db *attendanceConnection) GetFindByIdWeeksDetachment(week string, userid uint) (entity.Attendance, error) {
	var attendance entity.Attendance
	go func() {
		err := db.connection.Where("user_id=?", userid).
			Where("week_sub_detachment=?", week).
			Find(&attendance).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return attendance, err
}
