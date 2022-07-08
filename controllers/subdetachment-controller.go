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

type SubdetachmentController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	FindByIdDetachment(context *gin.Context)
}

type subdetachmentController struct {
	subDetachment InterfacesService.ISubDetachmentService
	jwt           InterfacesService.IJWTService
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

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.subDetachment.GetAllSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest,constantvariables.PERMISSION_DANIED))
}

// get
func (c *subdetachmentController) FindById(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.GetFindByIdSubDetachmentService(context)

		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest,constantvariables.PERMISSION_DANIED))
}

//FindByIdDetachment
func (c *subdetachmentController) FindByIdDetachment(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.GetFindByIdDetachmentSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest,constantvariables.PERMISSION_DANIED))
}

//create subdetachment method post
func (c *subdetachmentController) Create(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.SetCreateSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest,constantvariables.PERMISSION_DANIED))
}

//update subdetachment method put
func (c *subdetachmentController) Update(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.SetUpdateSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest,constantvariables.PERMISSION_DANIED))
}

// delete subdetachment
func (c *subdetachmentController) Remove(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.subDetachment.SetRemoveSubDetachmentService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest,constantvariables.PERMISSION_DANIED))

}
