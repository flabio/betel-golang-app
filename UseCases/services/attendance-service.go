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
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	attendances, err := attendanceService.IAttendance.GetAttendancesSubdetachment(uint(id), sub_detachment_id)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "ok", attendances))
}
func (attendanceService *attendanceService) AllAttendanceService(context *gin.Context) {
	attendances, err := attendanceService.IAttendance.GetAllAttendance()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "ok", attendances))
}
func (attendanceService *attendanceService) CreateAttendanceService(sub_detachment_id uint, context *gin.Context) {
	attendance := entity.Attendance{}
	var attendanceDTO dto.AttendanceDTO

	context.ShouldBind(&attendanceDTO)

	if validarAttendance(attendanceDTO, context, constantvariables.OPTION_CREATE) {
		return
	}
	attendanceDTO.SubDetachmentId = uint(sub_detachment_id)

	smapping.FillStruct(&attendance, smapping.MapFields(&attendanceDTO))
	existWeek, _ := attendanceService.IAttendance.GetFindByIdWeeksDetachment(attendanceDTO.WeekSubDetachment, attendanceDTO.UserId)

	if existWeek.Id > 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.WEEK_ALREADY)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return

	}
	data, errCreate := attendanceService.IAttendance.SetCreateAttendance(attendance)
	if errCreate != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errCreate.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, data))

}
func (attendanceService *attendanceService) UpdateAttendanceService(context *gin.Context) {
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
	attendanceById, _ := attendanceService.IAttendance.GetFindByIdAttendance(uint(attendanceDTO.Id))
	if attendanceById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, errUpdate := attendanceService.IAttendance.SetCreateAttendance(attendance)
	if errUpdate != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errUpdate.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, data))
}
func (attendanceService *attendanceService) RemoveAttendanceService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	attendanceById, errById := attendanceService.IAttendance.GetFindByIdAttendance(uint(id))
	if attendanceById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if errById != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errById.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, errRemove := attendanceService.IAttendance.SetRemoveAttendance(uint(id))
	if errRemove != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errRemove.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, data))
}

//weeks sud detachment
func (attendanceService *attendanceService) WeeksbySubDetachmentsAttendanceService(sub_detachment_id uint, context *gin.Context) {
	weeksSubDetachments, err := attendanceService.IWeeksDetachment.GetFindByIdWeeks(1)

	a := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)

	switch option {
	case 1:
		if len(a.WeekSubDetachment) == 0 || a.WeekSubDetachment == "" {
			res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.WEEK_REQUIRED)
			context.JSON(http.StatusBadRequest, res)
			return true
		}
		if a.UserId == 0 {
			context.JSON(http.StatusBadRequest, res)
			return true
		}
	case 2:
		if a.Id == 0 {
			context.JSON(http.StatusBadRequest, res)
			return true
		}
		if a.UserId == 0 {
			context.JSON(http.StatusBadRequest, res)
			return true
		}
	}

	return false
}
