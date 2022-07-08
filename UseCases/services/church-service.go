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

//create Service
func (churchService *churchService) CreateChurchService(context *gin.Context) {
	church := entity.Church{}
	var churchDTO dto.ChurchDTO

	errDTO := context.ShouldBind(&churchDTO)
	if errDTO != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errDTO.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := smapping.FillStruct(&church, smapping.MapFields(&churchDTO))
	checkError(err)

	data, err := churchService.IChurch.SetCreateChurch(church)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errDTO.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, data))
}

//update church
func (churchService *churchService) UpdateChurch(context *gin.Context) {

	church := entity.Church{}
	var churchDTO dto.ChurchDTO

	errDTO := context.ShouldBind(&churchDTO)
	if errDTO != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errDTO.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err := smapping.FillStruct(&church, smapping.MapFields(&churchDTO))
	checkError(err)

	findById, _ := churchService.IChurch.GetFindChurchById(uint(churchDTO.Id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := churchService.IChurch.SetCreateChurch(church)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, data))
}

//find by Id
func (churchService *churchService) FindChurchById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	church, err := churchService.IChurch.GetFindChurchById(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", church)
	context.JSON(http.StatusOK, res)
}

//list of chuch
func (churchService *churchService) AllChurch(context *gin.Context) {
	churchs, err := churchService.IChurch.GetAllChurch()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	results := utilities.BuildResponse(http.StatusOK, "OK!", churchs)
	context.JSON(http.StatusOK, results)

}

//delete church
func (churchService *churchService) DeleteChurch(context *gin.Context) {
	var church entity.Church
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	findById, _ := churchService.IChurch.GetFindChurchById(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.JSON(http.StatusBadRequest, res)
		return
	}
	status, err := churchService.IChurch.SetRemoveChurch(findById)
	if err != nil {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NOT_DELETED)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	if status {
		res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, church)
		context.JSON(http.StatusOK, res)
	}

}
