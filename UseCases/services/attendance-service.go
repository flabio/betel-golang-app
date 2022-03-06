package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type AttendanceService interface {
	CreateAttendanceService(sub_detachment_id uint, context *gin.Context)
	UpdateAttendanceService(context *gin.Context)
	RemoveAttendanceService(context *gin.Context)
	AllAttendanceService(context *gin.Context)
	AttendancesSubdetachmentService(sub_detachment_id uint, context *gin.Context)
	WeeksbySubDetachmentsAttendanceService(sub_detachment_id uint, context *gin.Context)
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
func (s *attendanceService) AttendancesSubdetachmentService(sub_detachment_id uint, context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	attendances, err := s.attendanceRepository.GetAttendancesSubdetachment(uint(id), sub_detachment_id)

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", attendances))
}
func (s *attendanceService) AllAttendanceService(context *gin.Context) {
	attendances, err := s.attendanceRepository.GetAllAttendance()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", attendances))
}
func (s *attendanceService) CreateAttendanceService(sub_detachment_id uint, context *gin.Context) {
	attendance := entity.Attendance{}
	var attendanceDTO dto.AttendanceDTO

	context.ShouldBind(&attendanceDTO)

	if validarAttendance(attendanceDTO, context, constantvariables.OPTION_CREATE) {
		return
	}
	attendanceDTO.SubDetachmentId = uint(sub_detachment_id)

	smapping.FillStruct(&attendance, smapping.MapFields(&attendanceDTO))
	existWeek, _ := s.attendanceRepository.GetFindByIdWeeksDetachment(attendanceDTO.WeekSubDetachment, attendanceDTO.UserId)

	if existWeek.Id > 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.ExtisByUserWeek(), context)
		return

	}
	data, errCreate := s.attendanceRepository.SetCreateAttendance(attendance)
	if errCreate != nil {
		validadErrors(errCreate, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))

}
func (s *attendanceService) UpdateAttendanceService(context *gin.Context) {
	attendance := entity.Attendance{}
	var attendanceDTO dto.AttendanceDTO

	context.ShouldBind(attendanceDTO)

	if validarAttendance(attendanceDTO, context, constantvariables.OPTION_EDIT) {
		return
	}

	errMap := smapping.FillStruct(&attendance, smapping.MapFields(&attendanceDTO))
	if errMap != nil {
		checkError(errMap)
		return
	}
	attendanceById, _ := s.attendanceRepository.GetFindByIdAttendance(uint(attendanceDTO.Id))
	if attendanceById.Id == 0 {
		validadErrorById(context)
		return
	}
	data, errUpdate := s.attendanceRepository.SetCreateAttendance(attendance)
	if errUpdate != nil {
		validadErrors(errUpdate, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(data))
}
func (s *attendanceService) RemoveAttendanceService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	attendanceById, errById := s.attendanceRepository.GetFindByIdAttendance(uint(id))
	if attendanceById.Id == 0 {
		validadErrorById(context)
		return
	}
	if errById != nil {
		validadErrors(errById, context)
		return
	}
	data, errRemove := s.attendanceRepository.SetRemoveAttendance(uint(id))
	if errRemove != nil {
		validadErrors(errRemove, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(data, attendanceById))
}

//weeks sud detachment
func (s *attendanceService) WeeksbySubDetachmentsAttendanceService(sub_detachment_id uint, context *gin.Context) {
	weeksSubDetachments, err := s.weekSubDetachmentRepository.GetFindByIdWeeksDetachment(1)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", weeksSubDetachments))
}

//validarAttendance
func validarAttendance(a dto.AttendanceDTO, context *gin.Context, option int) bool {
	context.ShouldBind(&a)

	switch option {
	case 1:
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
	case 2:
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
	}

	return false
}
