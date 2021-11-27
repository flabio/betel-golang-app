package controllers

import (
	"bete/Infrastructure/middleware"
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
	module services.ModuleService
	jwt    services.JWTService
}

func NewModuleController() ModuleController {

	return &moduleController{
		module: services.NewModuleService(),
		jwt:    services.NewJWTService(),
	}
}

//GET /lists of modules

func (c *moduleController) All(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.AllModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//GET /lists of modules

func (c *moduleController) ByRoleModule(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.AllByRoleModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//GET/module/1
// find module by id
func (c *moduleController) FindModuleById(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.FindModuleById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//POST/module
//create rol metho post
func (c *moduleController) Create(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.CreateModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//PUT/Module
//update module method put
func (c *moduleController) Update(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.UpdateModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//DELETE/module/1
// delete module
func (c *moduleController) Delete(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.DeleteModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//POST/rolemodule
//create rolemodule method post
func (c *moduleController) AddModuleRole(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.AddModuleRole(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//DELETE/rolemodule
//delete rolemodule method delete
func (c *moduleController) DeleteModuleRole(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.module.DeleteRoleModule(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
