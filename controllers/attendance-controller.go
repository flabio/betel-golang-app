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
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.attendance.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//Update
func (c *attendanceController) Update(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.attendance.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//Remove
func (c *attendanceController) Remove(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.attendance.Remove(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//All
func (c *attendanceController) All(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.attendance.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
