package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubdetachmentController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	FindByIdDetachment(context *gin.Context)
}

type subdetachmentController struct {
	subDetachment services.SubDetachmentService
	jwt           services.JWTService
}

func NewSubdetachmentController() SubdetachmentController {

	return &subdetachmentController{
		subDetachment: services.NewSubDetachmentService(),
		jwt:           services.NewJWTService(),
	}
}

//GET /subdetachment
// get list of subdetachment
func (c *subdetachmentController) All(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.subDetachment.GetAllSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// get
func (c *subdetachmentController) FindById(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.GetFindByIdSubDetachmentService(context)

		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//FindByIdDetachment
func (c *subdetachmentController) FindByIdDetachment(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.GetFindByIdDetachmentSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//create subdetachment method post
func (c *subdetachmentController) Create(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.SetCreateSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//update subdetachment method put
func (c *subdetachmentController) Update(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.SetUpdateSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// delete subdetachment
func (c *subdetachmentController) Remove(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.SetRemoveSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}
