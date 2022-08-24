package Interfaces

import "bete/Core/entity"

type IAttendance interface {
	SetCreateAttendance(attendance entity.Attendance) (entity.Attendance, error)
	SetUpdateAttendance(attendance entity.Attendance, Id uint) (entity.Attendance, error)
	SetRemoveAttendance(Id uint) (bool, error)
	GetFindByIdAttendance(Id uint) (entity.Attendance, error)
	GetFindByIdWeeksDetachment(Week string, IdUser uint) (entity.Attendance, error)
	GetAllAttendance() ([]entity.Attendance, error)
	GetAttendancesSubdetachment(IdUser uint, IdSubDetachment uint) ([]entity.Attendance, error)
}
