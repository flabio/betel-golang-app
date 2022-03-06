package controllers

import (
	"bete/Infrastructure/middleware"
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
	jwt        services.JWTService
	attendance services.AttendanceService
}

func NewAttendanceController() AttendanceController {
	return &attendanceController{
		jwt:        services.NewJWTService(),
		attendance: services.NewAttendanceService(),
	}
}

//Create
func (c *attendanceController) Create(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.CreateAttendanceService(claim.Subdetachmentid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//Update
func (c *attendanceController) Update(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.UpdateAttendanceService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//Remove
func (c *attendanceController) Remove(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.RemoveAttendanceService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//All
func (c *attendanceController) All(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.attendance.AllAttendanceService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//Attendances by subdetachment
func (c *attendanceController) AttendancesSubdetachment(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.AttendancesSubdetachmentService(claim.Subdetachmentid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//WeeksbySubDetachments
func (c *attendanceController) WeeksbySubDetachments(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.attendance.WeeksbySubDetachmentsAttendanceService(claim.Subdetachmentid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
