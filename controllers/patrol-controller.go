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

type PatrolController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type patrolController struct {
	patrol InterfacesService.IPatrolService
	jwt    InterfacesService.IJWTService
}

func NewPatrolController() PatrolController {

	return &patrolController{
		patrol: services.NewPatrolService(),
		jwt:    services.NewJWTService(),
	}
}

//GET /rols
// get list of rol
func (c *patrolController) All(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

// get
func (c *patrolController) FindById(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.FindById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//create rol metho post
func (c *patrolController) Create(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//update rol method put
func (c *patrolController) Update(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

// delete rol
func (c *patrolController) Remove(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.Remove(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))

}
