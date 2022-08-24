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

type ParentController interface {
	All(context *gin.Context)
	AllParentScout(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type parentController struct {
	parent InterfacesService.IParentService
	jwt    InterfacesService.IJWTService
}

func NewParentController() ParentController {
	return &parentController{
		parent: services.NewParentService(),
		jwt:    services.NewJWTService(),
	}
}

// All the parent
func (c *parentController) All(context *gin.Context) {
	rol := middleware.ValidadToken(c.jwt, context)
	if rol.Rol > 0 {
		c.parent.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// All the parent
func (c *parentController) AllParentScout(context *gin.Context) {
	rol := middleware.ValidadToken(c.jwt, context)
	if rol.Rol > 0 {
		c.parent.AllParentScout(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// Create parent
func (c *parentController) Create(context *gin.Context) {
	rol := middleware.ValidadToken(c.jwt, context)
	if rol.Rol > 0 {
		c.parent.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// Update parent
func (c *parentController) Update(context *gin.Context) {
	rol := middleware.ValidadToken(c.jwt, context)
	if rol.Rol > 0 {
		c.parent.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// Remove parent
func (c *parentController) Remove(context *gin.Context) {
	rol := middleware.ValidadToken(c.jwt, context)
	if rol.Rol > 0 {
		c.parent.Remove(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}
