package services

import (
	"bete/Core/repositorys"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CityService interface {
	All(context *gin.Context)
}

type cityService struct {
	cityRepository repositorys.CityRepository
}

//NewCityService
func NewCityService() CityService {
	return &cityService{
		cityRepository: repositorys.NewCityRepository(),
	}
}

//All
func (c *cityService) All(context *gin.Context) {
	var cities, err = c.cityRepository.GetAllCity()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", cities))
}
