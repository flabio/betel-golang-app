package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Interface DetachmenteController si a....

type DetachmentController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}

type detachmentController struct {
	detachmentService services.DetachmentService
	jwt               services.JWTService
}

func NewDetachmentController() DetachmentController {

	return &detachmentController{
		detachmentService: services.NewDetachmentService(),
		jwt:               services.NewJWTService(),
	}
}

//method:GET
//api/getAll
func (c *detachmentController) All(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//method:GET
//api/getFindById/:id
func (c *detachmentController) FindById(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.FindById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//method:POST
//api/setCreate
func (c *detachmentController) Create(context *gin.Context) {
	c.detachmentService.Create(context)
}

//method:PUT
//api/setUpdate
func (c *detachmentController) Update(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// method:DELETE
// api/setDelete?:id
func (c *detachmentController) Delete(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.Delete(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
