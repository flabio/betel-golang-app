package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type ChurchService interface {
	CreateChurch(context *gin.Context)
	UpdateChurch(context *gin.Context)
	DeleteChurch(context *gin.Context)
	FindChurchById(context *gin.Context)
	AllChurch(context *gin.Context)
}

type churchService struct {
	churchRepository repositorys.ChurchRepository
}

func NewChurchService() ChurchService {

	var churchRepository = repositorys.NewChurchRepository()
	return &churchService{
		churchRepository: churchRepository,
	}
}


//create Service
func (churchService *churchService) CreateChurch(context *gin.Context) {
	church := entity.Church{}
	var churchDTO dto.ChurchDTO

	errDTO := context.ShouldBind(&churchDTO)
	if errDTO != nil {
		validadErrors(errDTO, context)
		return
	}

	err := smapping.FillStruct(&church, smapping.MapFields(&churchDTO))
	checkError(err)

	data, err := churchService.churchRepository.CreateChurch(church)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))
}

//update church
func (churchService *churchService) UpdateChurch(context *gin.Context) {

	church := entity.Church{}
	var churchDTO dto.ChurchDTO

	errDTO := context.ShouldBind(&churchDTO)
	if errDTO != nil {
		validadErrors(errDTO, context)
		return
	}
	err := smapping.FillStruct(&church, smapping.MapFields(&churchDTO))
	checkError(err)

	findById, _ := churchService.churchRepository.FindChurchById(uint(churchDTO.Id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	data, err := churchService.churchRepository.UpdateChurch(church)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(data))
}

//find by Id
func (churchService *churchService) FindChurchById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	church, err := churchService.churchRepository.FindChurchById(uint(id))
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", church)
	context.JSON(http.StatusOK, res)
}

//list of chuch
func (churchService *churchService) AllChurch(context *gin.Context) {
	churchs, err := churchService.churchRepository.AllChurch()
	if err != nil {
		validadErrors(err, context)
		return
	}
	results := utilities.BuildResponse(true, "OK!", churchs)
	context.JSON(http.StatusOK, results)
	return
}

//delete church
func (churchService *churchService) DeleteChurch(context *gin.Context) {
	var church entity.Church
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		response := utilities.BuildErrorByIdResponse()
		context.JSON(http.StatusBadRequest, response)
		return
	}
	findById, _ := churchService.churchRepository.FindChurchById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	status, err := churchService.churchRepository.DeleteChurch(findById)
	if err != nil {
		validadErrorRemove(church, context)
		return
	}
	res := utilities.BuildDeteleteResponse(status, church)
	context.JSON(http.StatusOK, res)

}
