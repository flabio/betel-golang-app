package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type churchService struct {
	IChurch Interfaces.IChurch
}

func NewChurchService() InterfacesService.IChurchService {
	return &churchService{
		IChurch: repositorys.GetChurchInstance(),
	}
}

// create Service
func (churchService *churchService) CreateChurchService(context *gin.Context) {
	var churchDTO dto.ChurchDTO
	church, _ := getMappingChurch(churchDTO, context)
	data, err := churchService.IChurch.SetCreateChurch(church)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreatedResponse(data))
}

// update church
func (churchService *churchService) UpdateChurch(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	var churchDTO dto.ChurchDTO
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	church, err := getMappingChurch(churchDTO, context)

	findById, _ := churchService.IChurch.GetFindChurchById(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	data, err := churchService.IChurch.SetUpdateChurch(church, uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(data))
}

// find by Id
func (churchService *churchService) FindChurchById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	church, err := churchService.IChurch.GetFindChurchById(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(church))
}

// list of chuch
func (churchService *churchService) AllChurch(context *gin.Context) {
	churchs, err := churchService.IChurch.GetAllChurch()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(churchs))
}

// delete church
func (churchService *churchService) DeleteChurch(context *gin.Context) {
	var church entity.Church
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	findById, _ := churchService.IChurch.GetFindChurchById(uint(id))
	if findById.Id == 0 {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	status, err := churchService.IChurch.SetRemoveChurch(findById)
	if err != nil {
		response := utilities.BuildErrResponse(constantvariables.NOT_DELETED)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	if status {
		context.JSON(http.StatusOK, utilities.BuildRemovedResponse(church))

	}

}

func getMappingChurch(churchDTO dto.ChurchDTO, context *gin.Context) (entity.Church, error) {
	var church entity.Church
	err := context.ShouldBind(&churchDTO)
	if err != nil {
		return church, err
	}
	err = smapping.FillStruct(&church, smapping.MapFields(&churchDTO))
	if err != nil {
		return church, err
	}
	return church, nil
}
