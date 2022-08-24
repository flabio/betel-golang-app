package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

// NewAttendanceRepository()
func GetAttendanceInstance() Interfaces.IAttendance {
	// sync.Once bloquear de manera
	var (
		_OPEN *OpenConnections
		_ONCE sync.Once
	)
	// Do aseguro que se ejecute una unica vez de manera seguro
	_ONCE.Do(func() {
		_OPEN = &OpenConnections{

			connection: entity.Factory(constantvariables.OPTION_FACTORY_DB),
		}
	})
	return _OPEN
}

func (db *OpenConnections) SetCreateAttendance(attendance entity.Attendance) (entity.Attendance, error) {
	db.mux.Lock()
	err := db.connection.Save(&attendance).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return attendance, err
}

func (db *OpenConnections) SetUpdateAttendance(attendance entity.Attendance, Id uint) (entity.Attendance, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", Id).Save(&attendance).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return attendance, err
}

/*
@param Id is Attendace ,of type uint
*/
func (db *OpenConnections) SetRemoveAttendance(Id uint) (bool, error) {
	var attendance entity.Attendance
	db.mux.Lock()
	err := db.connection.Delete(&attendance, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err != nil {
		return false, err
	}
	return true, err
}
func (db *OpenConnections) GetAllAttendance() ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	db.mux.Lock()
	err := db.connection.Preload("User").Preload("SubDetachment").Find(&attendance).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return attendance, err
}

/*
@param IdUser is the user, of type uint
@param IdSubDetachment is the Subdetachment, of type uint
*/
func (db *OpenConnections) GetAttendancesSubdetachment(IdUser uint, IdSubDetachment uint) ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	db.mux.Lock()
	err := db.connection.Preload("User").
		Preload("SubDetachment").
		Where("user_id=?", IdUser).
		Where("sub_detachment_id=?", IdSubDetachment).
		Find(&attendance).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return attendance, err
}

/*
@param Id is the attendance, of type uint
*/
func (db *OpenConnections) GetFindByIdAttendance(Id uint) (entity.Attendance, error) {
	var attendance entity.Attendance
	db.mux.Lock()
	err := db.connection.Find(&attendance, Id).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return attendance, err
}

/*
@param Week is the attendance, of type uint
@param IdUser is the attendance, of type uint
*/
func (db *OpenConnections) GetFindByIdWeeksDetachment(Week string, IdUser uint) (entity.Attendance, error) {
	var attendance entity.Attendance
	db.mux.Lock()
	err := db.connection.Where("user_id=?", IdUser).
		Where("week_sub_detachment=?", Week).
		Find(&attendance).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return attendance, err
}
