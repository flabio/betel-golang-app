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

// patrolService
type patrolService struct {
	IPatrol Interfaces.IPatrol
}

func NewPatrolService() InterfacesService.IPatrolService {
	return &patrolService{
		IPatrol: repositorys.NewPatrolRepository(),
	}
}

// Create
func (patrolService *patrolService) Create(context *gin.Context) {
	var dto dto.PatrolDTO
	patrol, msg := getMappingPatrol(dto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	filename, err := UploadFile(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	patrol.Archives = filename

	res, err := patrolService.IPatrol.SetCreatePatrol(patrol)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildCreatedResponse(res))
}

// Update
func (patrolService *patrolService) Update(context *gin.Context) {
	var dto dto.PatrolDTO
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	patrol, msg := getMappingPatrol(dto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}

	findById, _ := patrolService.IPatrol.GetFindByIdPatrol(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	filename, err := UploadFile(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if len(findById.Archives) == 0 {
		patrol.Archives = filename
	} else {
		if filename != "" {
			patrol.Archives = filename
		} else {
			patrol.Archives = findById.Archives
		}
	}

	res, err := patrolService.IPatrol.SetUpdatePatrol(patrol, uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildResponse(res))
}

// Remove
func (patrolService *patrolService) Remove(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	findById, _ := patrolService.IPatrol.GetFindByIdPatrol(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	res, err := patrolService.IPatrol.SetRemovePatrol(findById.Id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildRemovedResponse(res))
}

// FindById
func (patrolService *patrolService) FindById(context *gin.Context) {
	var response []dto.PatrolResponse
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	data, _ := patrolService.IPatrol.GetFindByIdPatrol(uint(id))
	if data.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	response = append(response, patrolResponse(data))
	context.JSON(http.StatusOK, utilities.BuildResponse(response))
}

// All
func (patrolService *patrolService) All(context *gin.Context) {
	var patrolList []dto.PatrolResponse
	res, err := patrolService.IPatrol.GetAllPatrol()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, data := range res {
		patrolList = append(patrolList, patrolResponse(data))
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(patrolList))
}

// validate
func getMappingPatrol(dto dto.PatrolDTO, context *gin.Context) (entity.Patrol, string) {
	patrol := entity.Patrol{}
	err := context.ShouldBind(&dto)
	if err != nil {
		return patrol, utilities.GetMsgErrorRequired(err)
	}
	err = smapping.FillStruct(&patrol, smapping.MapFields(&dto))
	if err != nil {
		return patrol, err.Error()
	}
	return patrol, ""
}

func patrolResponse(data entity.Patrol) dto.PatrolResponse {
	return dto.PatrolResponse{
		Id:                data.Id,
		Name:              data.Name,
		Archives:          data.Archives,
		Active:            data.Active,
		SubDetachmentId:   data.SubDetachmentId,
		SubDetachmentName: data.SubDetachment.Name,
	}
}
