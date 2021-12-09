package controllers

import (
	"bete/Infrastructure/middleware"
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
	patrol services.PatrolService
	jwt    services.JWTService
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

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// get
func (c *patrolController) FindById(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.FindById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//create rol metho post
func (c *patrolController) Create(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//update rol method put
func (c *patrolController) Update(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// delete rol
func (c *patrolController) Remove(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.patrol.Remove(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}
