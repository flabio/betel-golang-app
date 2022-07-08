package controllers

import (
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/Infrastructure/middleware"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttendanceController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	All(context *gin.Context)
	AttendancesSubdetachment(context *gin.Context)
	WeeksbySubDetachments(context *gin.Context)
}

type attendanceController struct {
	jwt        InterfacesService.IJWTService
	attendance InterfacesService.IAttendanceService
}

func NewAttendanceController() AttendanceController {
	return &attendanceController{
		jwt:        services.NewJWTService(),
		attendance: services.NewAttendanceService(),
	}
}

//Create
func (c *attendanceController) Create(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.CreateAttendanceService(claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//Update
func (c *attendanceController) Update(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.UpdateAttendanceService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//Remove
func (c *attendanceController) Remove(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.RemoveAttendanceService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//All
func (c *attendanceController) All(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.attendance.AllAttendanceService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//Attendances by subdetachment
func (c *attendanceController) AttendancesSubdetachment(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.AttendancesSubdetachmentService(claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//WeeksbySubDetachments
func (c *attendanceController) WeeksbySubDetachments(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.WeeksbySubDetachmentsAttendanceService(claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}
