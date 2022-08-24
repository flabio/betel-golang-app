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

type visitService struct {
	IVisit Interfaces.IVisit
}

func NewVisitService() InterfacesService.IVisitService {
	return &visitService{
		IVisit: repositorys.GetVisitInstance(),
	}
}

// Create visit
func (visitService *visitService) SetCreateVisitService(context *gin.Context) {
	var visitDto dto.VisitDTO
	var visit entity.Visit
	context.ShouldBind(&visitDto)
	if validarVisit(visitDto, context, constantvariables.OPTION_CREATE) {
		return
	}
	smapping.FillStruct(&visit, smapping.MapFields(&visitDto))
	data, err := visitService.IVisit.SetCreateVisit(visit)

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildCreatedResponse(data))
}

// Update
func (visitService *visitService) SetUpdateVisitService(context *gin.Context) {
	var visitDto dto.VisitDTO
	var visit entity.Visit
	context.ShouldBind(&visitDto)
	if validarVisit(visitDto, context, constantvariables.OPTION_EDIT) {
		return
	}
	smapping.FillStruct(&visit, smapping.MapFields(&visitDto))
	existVisit, _ := visitService.IVisit.GetFindByIdVisit(visitDto.Id)
	if existVisit.Id == 0 {
		res := utilities.BuildErrResponse(constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := visitService.IVisit.SetCreateVisit(visit)

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(data))
}

// All visit
func (visitService *visitService) GetAllVisitService(context *gin.Context) {
	data, err := visitService.IVisit.GetAllVisit()

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(data))
}

// AllVisitByUser by iduser and idsubdetachment
func (visitService *visitService) GetAllVisitByUserVisitService(subDetachmentId uint, context *gin.Context) {
	id, errId := strconv.ParseInt(context.Param("id"), 0, 0)
	if errId != nil {
		res := utilities.BuildErrResponse(errId.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := visitService.IVisit.GetAllVisitByUserVisit(uint(id), uint(subDetachmentId))

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(data))

}

func (visitService *visitService) SetRemoveVisitService(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	existVisit, _ := visitService.IVisit.GetFindByIdVisit(uint(id))
	if existVisit.Id == 0 {
		res := utilities.BuildErrResponse(constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	parent, err := visitService.IVisit.SetRemoveVisit(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(constantvariables.NOT_DELETED)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if parent {
		context.JSON(http.StatusOK, utilities.BuildRemovedResponse(existVisit))

	}
}

// validarVisiCreate
func validarVisit(visit dto.VisitDTO, context *gin.Context, option int) bool {
	context.ShouldBind(&visit)
	switch option {
	case 1:
		if validarVisitField(visit, context) {
			return true
		}
	case 2:

		if visit.Id == 0 {
			res := utilities.BuildErrResponse(constantvariables.ID)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return true
		}
		if validarVisitField(visit, context) {
			return true
		}
	}

	return false
}

// validarVisiEdit
func validarVisitField(visit dto.VisitDTO, context *gin.Context) bool {
	context.ShouldBind(&visit)
	if len(visit.State) == 0 || visit.State == "" {
		res := utilities.BuildErrResponse(constantvariables.STATE)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	if len(visit.Description) == 0 || visit.Description == "" {
		res := utilities.BuildErrResponse(constantvariables.DESCRIPTION)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	if visit.UserId == 0 {

		res := utilities.BuildErrResponse(constantvariables.USER_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	if visit.SubDetachmentId == 0 {
		res := utilities.BuildErrResponse(constantvariables.SUB_DETACHMENT_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	return false
}
