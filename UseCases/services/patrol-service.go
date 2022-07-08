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

//patrolService
type patrolService struct {
	IPatrol Interfaces.IPatrol
}

func NewPatrolService() InterfacesService.IPatrolService {
	return &patrolService{
		IPatrol: repositorys.NewPatrolRepository(),
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
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	patrol.Archives = filename

	res, err := patrolService.IPatrol.SetCreatePatrol(patrol)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, res))
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

	findById, _ := patrolService.IPatrol.GetFindByIdPatrol(uint(dto.Id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	filename, err := UploadFile(context)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
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

	res, err := patrolService.IPatrol.SetCreatePatrol(patrol)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, res))
}

//Remove
func (patrolService *patrolService) Remove(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	findById, _ := patrolService.IPatrol.GetFindByIdPatrol(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res, err := patrolService.IPatrol.SetRemovePatrol(findById.Id)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, res))
}

//FindById
func (patrolService *patrolService) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	findById, _ := patrolService.IPatrol.GetFindByIdPatrol(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "ok", findById))
}

//All
func (patrolService *patrolService) All(context *gin.Context) {
	res, err := patrolService.IPatrol.GetAllPatrol()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "ok", res))
}

//validate
func validatePatroCreate(dto dto.PatrolDTO, context *gin.Context) bool {
	context.ShouldBind(&dto)
	if len(dto.Name) == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NAME)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	if dto.SubDetachmentId == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	return false
}
