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

type ModuleController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	All(context *gin.Context)
	ByRoleModule(context *gin.Context)
	FindModuleById(context *gin.Context)
	Delete(context *gin.Context)
	AddModuleRole(context *gin.Context)
	DeleteModuleRole(context *gin.Context)
}

type moduleController struct {
	module InterfacesService.IModuleService
	jwt    InterfacesService.IJWTService
}

func NewModuleController() ModuleController {

	return &moduleController{
		module: services.NewModuleService(),
		jwt:    services.NewJWTService(),
	}
}

//GET /lists of modules

func (c *moduleController) All(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.module.AllModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//GET /lists of modules

func (c *moduleController) ByRoleModule(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.module.AllByRoleModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//GET/module/1
// find module by id
func (c *moduleController) FindModuleById(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.module.FindModuleById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//POST/module
//create rol metho post
func (c *moduleController) Create(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.module.CreateModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//PUT/Module
//update module method put
func (c *moduleController) Update(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.module.UpdateModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//DELETE/module/1
// delete module
func (c *moduleController) Delete(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.module.DeleteModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//POST/rolemodule
//create rolemodule method post
func (c *moduleController) AddModuleRole(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.module.AddModuleRole(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//DELETE/rolemodule
//delete rolemodule method delete
func (c *moduleController) DeleteModuleRole(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.module.DeleteRoleModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}
