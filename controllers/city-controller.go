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

type CityController interface {
	All(context *gin.Context)
}
type cityController struct {
	city InterfacesService.ICityService
	jwt  InterfacesService.IJWTService
}

func NewCityController() CityController {
	return &cityController{
		city: services.NewCityService(),
		jwt:  services.NewJWTService(),
	}
}

//All is method GET
func (c *cityController) All(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.city.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}
