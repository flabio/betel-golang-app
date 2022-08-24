package services

import (
	"bete/Core/Interfaces"
	"bete/Core/repositorys"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cityService struct {
	ICity Interfaces.ICity
}

// NewCityService
func NewCityService() InterfacesService.ICityService {
	return &cityService{
		ICity: repositorys.GetCityInstance(),
	}
}

// All
func (cityService *cityService) All(context *gin.Context) {
	var cities, err = cityService.ICity.GetAllCity()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(cities))
}
