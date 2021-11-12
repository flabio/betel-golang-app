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

//PatrolService
type PatrolService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}

//patrolService
type patrolService struct {
	patrolRepository repositorys.PatrolRepository
}

func NewPatrolService() PatrolService {
	patrolRepository := repositorys.NewPatrolRepository()
	return &patrolService{
		patrolRepository: patrolRepository,
	}
}

//Create
func (patrolService *patrolService) Create(context *gin.Context) {
	patrol := entity.Patrol{}
	var dto dto.PatrolDTO
	context.ShouldBind(&dto)
	if validatePatroCreate(dto, context) {
		return
	}
	smapping.FillStruct(&patrol, smapping.MapFields(&dto))

	filename, err := UploadFile(context)
	patrol.Archives = filename

	res, err := patrolService.patrolRepository.Create(patrol)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(res))
}

//Update
func (patrolService *patrolService) Update(context *gin.Context) {
	patrol := entity.Patrol{}
	var dto dto.PatrolDTO
	context.ShouldBind(&dto)
	if validatePatroCreate(dto, context) {
		return
	}
	smapping.FillStruct(&patrol, smapping.MapFields(&dto))

	findById, _ := patrolService.patrolRepository.FindById(uint(dto.Id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	filename, err := UploadFile(context)

	if len(findById.Archives) == 0 {
		patrol.Archives = filename
	} else {
		if filename != "" {
			patrol.Archives = filename
		} else {
			patrol.Archives = findById.Archives
		}
	}

	res, err := patrolService.patrolRepository.Update(patrol)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(res))
}

//Remove
func (patrolService *patrolService) Remove(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	findById, _ := patrolService.patrolRepository.FindById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	res, err := patrolService.patrolRepository.Remove(findById.Id)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(res, findById))
}

//FindById
func (patrolService *patrolService) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	findById, _ := patrolService.patrolRepository.FindById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", findById))
}

//All
func (patrolService *patrolService) All(context *gin.Context) {
	res, err := patrolService.patrolRepository.All()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", res))
}

//validate
func validatePatroCreate(dto dto.PatrolDTO, context *gin.Context) bool {
	context.ShouldBind(&dto)
	if len(dto.Name) == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}
	if dto.SubDetachmentId == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredSubDetachment(), context)
		return true
	}
	return false
}
