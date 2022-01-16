package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VisitController interface {
	All(context *gin.Context)
	AllVisitByUserAndSubDatachment(context *gin.Context)
	CreateVisit(context *gin.Context)
	UpdateVisit(context *gin.Context)
	RemoveVisit(context *gin.Context)
}

type visitController struct {
	visit services.VisitService
	jwt   services.JWTService
}

func NewVisitController() VisitController {
	return &visitController{
		visit: services.NewVisitService(),
		jwt:   services.NewJWTService(),
	}
}

//All the visit
func (c *visitController) All(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.visit.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}
func (c *visitController) AllVisitByUserAndSubDatachment(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.visit.AllVisitByUser(claim.Subdetachmentid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
func (c *visitController) CreateVisit(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.visit.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
func (c *visitController) UpdateVisit(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.visit.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
func (c *visitController) RemoveVisit(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.visit.Remove(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
