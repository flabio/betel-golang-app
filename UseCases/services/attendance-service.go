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
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	All(context *gin.Context)
}

type attendanceService struct {
	attendanceRepository repositorys.AttendanceRepository
}

func NewAttendanceService() AttendanceService {
	return &attendanceService{
		attendanceRepository: repositorys.NewAttendanceRepository(),
	}
}
func (s *attendanceService) All(context *gin.Context) {
	attendances, err := s.attendanceRepository.All()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", attendances))
}
func (s *attendanceService) Create(context *gin.Context) {
	attendance := entity.Attendance{}
	var attendanceDTO dto.AttendanceDTO

	context.ShouldBind(attendanceDTO)

	if validarAttendanceCreate(attendanceDTO, context) {
		return
	}
	err := smapping.FillStruct(&attendance, smapping.MapFields(&attendanceDTO))
	if err != nil {
		checkError(err)
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

//validarAttendanceCreate
func validarAttendanceCreate(a dto.AttendanceDTO, context *gin.Context) bool {
	context.ShouldBind(&a)
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
