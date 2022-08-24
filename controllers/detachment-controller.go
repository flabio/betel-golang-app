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

// Interface DetachmenteController si a....

type DetachmentController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}

type detachmentController struct {
	detachmentService InterfacesService.IDetachmentService
	jwt               InterfacesService.IJWTService
}

func NewDetachmentController() DetachmentController {

	return &detachmentController{
		detachmentService: services.NewDetachmentService(),
		jwt:               services.NewJWTService(),
	}
}

// method:GET
// api/getAll
func (c *detachmentController) All(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// method:GET
// api/getFindById/:id
func (c *detachmentController) FindById(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.FindById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// method:POST
// api/setCreate
func (c *detachmentController) Create(context *gin.Context) {
	c.detachmentService.Create(context)
}

// method:PUT
// api/setUpdate
func (c *detachmentController) Update(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// method:DELETE
// api/setDelete?:id
func (c *detachmentController) Delete(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.detachmentService.Delete(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}
