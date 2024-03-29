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

type ChurchController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}

type churchController struct {
	church InterfacesService.IChurchService
	jwt    InterfacesService.IJWTService
}

func NewChurchController() ChurchController {

	return &churchController{
		church: services.NewChurchService(),
		jwt:    services.NewJWTService(),
	}
}

// method:GET
// api/getAll
func (c *churchController) All(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.church.AllChurch(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// method:GET
// api/getFindById/:id
func (c *churchController) FindById(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.church.FindChurchById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// method:POST
// api/setCreate
func (c *churchController) Create(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.church.CreateChurchService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))

}

// method:PUT
// api/setUpdate
func (c *churchController) Update(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.church.UpdateChurch(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// method:DELETE
// api/setDelete?:id
func (c *churchController) Delete(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.church.DeleteChurch(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}
