package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CityController interface {
	All(context *gin.Context)
}
type cityController struct {
	city services.CityService
	jwt  services.JWTService
}

func NewCityController() CityController {
	return &cityController{
		city: services.NewCityService(),
		jwt:  services.NewJWTService(),
	}
}

//All is method GET
func (c *cityController) All(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.city.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
