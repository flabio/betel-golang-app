package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type AttendanceService interface {
	Create(sub_detachment_id uint, context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	All(context *gin.Context)
	AttendancesSubdetachment(sub_detachment_id uint, context *gin.Context)
	WeeksbySubDetachments(sub_detachment_id uint, context *gin.Context)
}

type attendanceService struct {
	attendanceRepository        repositorys.AttendanceRepository
	weekSubDetachmentRepository repositorys.WeeksDetachmentRepository
}

func NewAttendanceService() AttendanceService {
	return &attendanceService{
		attendanceRepository:        repositorys.NewAttendanceRepository(),
		weekSubDetachmentRepository: repositorys.NewWeeksDetachmentRepository(),
	}
}
func (s *attendanceService) AttendancesSubdetachment(sub_detachment_id uint, context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	attendances, err := s.attendanceRepository.AttendancesSubdetachment(uint(id), sub_detachment_id)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", attendances))
}
func (s *attendanceService) All(context *gin.Context) {
	attendances, err := s.attendanceRepository.All()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", attendances))
}
func (s *attendanceService) Create(sub_detachment_id uint, context *gin.Context) {
	attendance := entity.Attendance{}
	var attendanceDTO dto.AttendanceDTO

	context.ShouldBind(&attendanceDTO)

	if validarAttendanceCreate(attendanceDTO, context) {
		return
	}
	attendanceDTO.SubDetachmentId = uint(sub_detachment_id)

	smapping.FillStruct(&attendance, smapping.MapFields(&attendanceDTO))
	existWeek, _ := s.attendanceRepository.FindByIdWeeksDetachment(attendanceDTO.WeekSubDetachment, attendanceDTO.UserId)

	if existWeek.Id > 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.ExtisByUserWeek(), context)
		return

	}
	data, errCreate := s.attendanceRepository.Create(attendance)
	if errCreate != nil {
		validadErrors(errCreate, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))

}
func (s *attendanceService) Update(context *gin.Context) {
	attendance := entity.Attendance{}
	var attendanceDTO dto.AttendanceDTO

	context.ShouldBind(attendanceDTO)

	if validarAttendanceEdit(attendanceDTO, context) {
		return
	}

	errMap := smapping.FillStruct(&attendance, smapping.MapFields(&attendanceDTO))
	if errMap != nil {
		checkError(errMap)
		return
	}
	attendanceById, _ := s.attendanceRepository.FindByIdAttendance(uint(attendanceDTO.Id))
	if attendanceById.Id == 0 {
		validadErrorById(context)
		return
	}
	data, errUpdate := s.attendanceRepository.Update(attendance)
	if errUpdate != nil {
		validadErrors(errUpdate, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(data))
}
func (s *attendanceService) Remove(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	attendanceById, errById := s.attendanceRepository.FindByIdAttendance(uint(id))
	if attendanceById.Id == 0 {
		validadErrorById(context)
		return
	}
	if errById != nil {
		validadErrors(errById, context)
		return
	}
	data, errRemove := s.attendanceRepository.Remove(uint(id))
	if errRemove != nil {
		validadErrors(errRemove, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(data, attendanceById))
}

//weeks sud detachment
func (s *attendanceService) WeeksbySubDetachments(sub_detachment_id uint, context *gin.Context) {
	weeksSubDetachments, err := s.weekSubDetachmentRepository.FindByIdWeeksDetachment(1)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", weeksSubDetachments))
}

//validarAttendanceCreate
func validarAttendanceCreate(a dto.AttendanceDTO, context *gin.Context) bool {
	context.ShouldBind(&a)
	if len(a.WeekSubDetachment) == 0 || a.WeekSubDetachment == "" {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredWeek(), context)
		return true
	}
	if a.UserId == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}

	return false
}

//validarAttendanceEdit
func validarAttendanceEdit(a dto.AttendanceDTO, context *gin.Context) bool {
	context.ShouldBind(&a)
	if a.Id == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}
	if a.UserId == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}

	return false
}
