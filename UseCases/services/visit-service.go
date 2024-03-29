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
	var visitDto dto.VisitRequest

	visit, msg := getMappingVisit(visitDto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	data, err := visitService.IVisit.SetCreateVisit(visit)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildCreatedResponse(data))
}

// Update
func (visitService *visitService) SetUpdateVisitService(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	var visitDto dto.VisitRequest

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	visit, msg := getMappingVisit(visitDto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	existVisit, _ := visitService.IVisit.GetFindByIdVisit(visitDto.Id)
	if existVisit.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	data, err := visitService.IVisit.SetCreateVisit(visit)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildUpdatedResponse(data))
}

// All visit
func (visitService *visitService) GetAllVisitService(context *gin.Context) {
	var visitLists []dto.VisitResponse
	data, err := visitService.IVisit.GetAllVisit()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, item := range data {
		visit := getVisitResponse(item)
		visitLists = append(visitLists, visit)
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(visitLists))
}

// AllVisitByUser by iduser and idsubdetachment
func (visitService *visitService) GetAllVisitByUserVisitService(subDetachmentId uint, context *gin.Context) {
	var visitLists []dto.VisitResponse
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	data, err := visitService.IVisit.GetAllVisitByUserVisit(uint(id), uint(subDetachmentId))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, item := range data {
		visit := getVisitResponse(item)
		visitLists = append(visitLists, visit)
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(visitLists))

}

func (visitService *visitService) SetRemoveVisitService(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	existVisit, _ := visitService.IVisit.GetFindByIdVisit(uint(id))
	if existVisit.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	parent, err := visitService.IVisit.SetRemoveVisit(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.NOT_DELETED))
		return
	}
	if parent {
		context.JSON(http.StatusCreated, utilities.BuildRemovedResponse(existVisit))

	}
}

func getMappingVisit(visitDto dto.VisitRequest, context *gin.Context) (entity.Visit, string) {
	visit := entity.Visit{}
	err := context.ShouldBind(&visitDto)
	if err != nil {
		return visit, utilities.GetMsgErrorRequired(err)
	}
	err = smapping.FillStruct(&visit, smapping.MapFields(&visitDto))
	if err != nil {
		return visit, err.Error()
	}
	return visit, ""
}

func getVisitResponse(data entity.Visit) dto.VisitResponse {
	visit := dto.VisitResponse{
		Id:                data.Id,
		State:             data.State,
		Description:       data.Description,
		UserFullName:      data.User.Name + " " + data.User.LastName,
		UserId:            data.UserId,
		SubDetachmentName: data.SubDetachment.Name,
		SubDetachmentId:   data.SubDetachmentId,
		Active:            data.Active,
	}
	return visit
}
