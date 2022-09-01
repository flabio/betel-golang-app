package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type attendanceService struct {
	IAttendance      Interfaces.IAttendance
	IWeeksDetachment Interfaces.IWeeksDetachment
}

func NewAttendanceService() InterfacesService.IAttendanceService {
	return &attendanceService{
		IAttendance:      repositorys.GetAttendanceInstance(),
		IWeeksDetachment: repositorys.NewWeeksDetachmentRepository(),
	}
}

func (attendanceService *attendanceService) AttendancesSubdetachmentService(sub_detachment_id uint, context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	attendances, err := attendanceService.IAttendance.GetAttendancesSubdetachment(uint(id), sub_detachment_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(attendances))
}
func (attendanceService *attendanceService) AllAttendanceService(context *gin.Context) {
	attendances, err := attendanceService.IAttendance.GetAllAttendance()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(attendances))
}
func (attendanceService *attendanceService) CreateAttendanceService(sub_detachment_id uint, context *gin.Context) {

	var attendanceDTO dto.AttendanceDTO
	attendanceDTO.SubDetachmentId = uint(sub_detachment_id)

	attendance, msg := getMappingAttendance(attendanceDTO, sub_detachment_id, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	existWeek, _ := attendanceService.IAttendance.GetFindByIdWeeksDetachment(attendanceDTO.WeekSubDetachment, attendanceDTO.UserId)

	if existWeek.Id > 0 {
		res := utilities.BuildErrResponse(constantvariables.WEEK_ALREADY)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return

	}
	data, err := attendanceService.IAttendance.SetCreateAttendance(attendance)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreatedResponse(data))

}
func (attendanceService *attendanceService) UpdateAttendanceService(context *gin.Context) {
	var attendanceDTO dto.AttendanceDTO
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	attendance, msg := getMappingAttendance(attendanceDTO, attendanceDTO.SubDetachmentId, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	attendanceById, _ := attendanceService.IAttendance.GetFindByIdAttendance(uint(id))
	if attendanceById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	data, err := attendanceService.IAttendance.SetUpdateAttendance(attendance, uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(data))
}
func (attendanceService *attendanceService) RemoveAttendanceService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	attendanceById, err := attendanceService.IAttendance.GetFindByIdAttendance(uint(id))
	if attendanceById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	data, err := attendanceService.IAttendance.SetRemoveAttendance(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildRemovedResponse(data))
}

func getMappingAttendance(attendanceDTO dto.AttendanceDTO, sub_detachment_id uint, context *gin.Context) (entity.Attendance, string) {
	var attendance entity.Attendance
	err := context.ShouldBind(&attendanceDTO)
	if err != nil {
		msgErros := utilities.GetMsgErrorRequired(err)
		return attendance, msgErros
	}
	attendanceDTO.SubDetachmentId = sub_detachment_id

	err = smapping.FillStruct(&attendance, smapping.MapFields(&attendanceDTO))
	if err != nil {
		return attendance, err.Error()
	}
	return attendance, ""

}
