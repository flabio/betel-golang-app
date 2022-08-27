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

//RolController is a..

type RolController interface {
	All(context *gin.Context)
	AllGroupRol(context *gin.Context)
	AllRoleModule(context *gin.Context)
	FindRol(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type rolController struct {
	rol InterfacesService.IRolService
	jwt InterfacesService.IJWTService
}

func NewRolController() RolController {

	return &rolController{
		rol: services.NewRolService(),
		jwt: services.NewJWTService(),
	}
}

// GET /rols
// get list of rol
func (c *rolController) All(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.GetAllService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// GET /rols
// get list of rol
func (c *rolController) AllGroupRol(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.GetAllGroupRolService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// GET /role module
// get list of role module
func (c *rolController) AllRoleModule(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.GetAllRoleModuleService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// get
func (c *rolController) FindRol(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.GetFindByIdService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))

}

// create rol metho post
func (c *rolController) Create(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// update rol method put
func (c *rolController) Update(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.SetUpdateService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// delete rol
func (c *rolController) Remove(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.SetRemoveService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}
