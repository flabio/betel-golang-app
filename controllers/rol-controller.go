package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RolController is a..

type RolController interface {
	All(context *gin.Context)
	FindRol(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type rolController struct {
	rol services.RolService
	jwt services.JWTService
}

func NewRolController() RolController {

	return &rolController{
		rol: services.NewRolService(),
		jwt: services.NewJWTService(),
	}
}

//GET /rols
// get list of rol
func (c *rolController) All(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.rol.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// get
func (c *rolController) FindRol(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.rol.FindById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

//create rol metho post
func (c *rolController) Create(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.rol.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

//update rol method put
func (c *rolController) Update(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.rol.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

// delete rol
func (c *rolController) Delete(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.rol.Delete(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}
