package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type VisitService interface {
	SetCreateVisitService(context *gin.Context)
	SetUpdateVisitService(context *gin.Context)
	GetAllVisitService(context *gin.Context)
	GetAllVisitByUserVisitService(subDetachmentId uint, context *gin.Context)
	SetRemoveVisitService(context *gin.Context)
}

type visitService struct {
	visitRepository repositorys.VisitRepository
}

func NewVisitService() VisitService {
	return &visitService{
		visitRepository: repositorys.NewVisitConnection(),
	}
}

//Create visit
func (service *visitService) SetCreateVisitService(context *gin.Context) {
	var visitDto dto.VisitDTO
	var visit entity.Visit
	context.ShouldBind(&visitDto)
	if validarVisit(visitDto, context, constantvariables.OPTION_CREATE) {
		return
	}
	smapping.FillStruct(&visit, smapping.MapFields(&visitDto))
	data, err := service.visitRepository.SetCreateVisit(visit)

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))
}

//Update
func (service *visitService) SetUpdateVisitService(context *gin.Context) {
	var visitDto dto.VisitDTO
	var visit entity.Visit
	context.ShouldBind(&visitDto)
	if validarVisit(visitDto, context, constantvariables.OPTION_EDIT) {
		return
	}
	smapping.FillStruct(&visit, smapping.MapFields(&visitDto))
	existVisit, _ := service.visitRepository.GetFindByIdVisit(visitDto.Id)
	if existVisit.Id == 0 {
		validadErrorById(context)
		return
	}
	data, err := service.visitRepository.SetCreateVisit(visit)

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))
}

//All visit
func (service *visitService) GetAllVisitService(context *gin.Context) {
	data, err := service.visitRepository.GetAllVisit()

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", data))
}

//AllVisitByUser by iduser and idsubdetachment
func (service *visitService) GetAllVisitByUserVisitService(subDetachmentId uint, context *gin.Context) {
	id, errId := strconv.ParseInt(context.Param("id"), 0, 0)
	if errId != nil {
		validadErrors(errId, context)
		return
	}
	data, err := service.visitRepository.GetAllVisitByUserVisit(uint(id), uint(subDetachmentId))

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", data))

}

func (service *visitService) SetRemoveVisitService(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	existVisit, _ := service.visitRepository.GetFindByIdVisit(uint(id))
	if existVisit.Id == 0 {
		validadErrorById(context)
		return
	}
	parent, err := service.visitRepository.SetRemoveVisit(uint(id))
	if err != nil {
		validadErrorRemove(id, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(parent, existVisit))
}

//validarVisiCreate
func validarVisit(visit dto.VisitDTO, context *gin.Context, option int) bool {
	context.ShouldBind(&visit)
	msg := utilities.MessageRequired{}
	switch option {
	case 1:
		if validarVisitField(visit, context) {
			return true
		}
	case 2:

		if visit.Id == 0 {
			validadRequiredMsg(msg.RequiredId(), context)
			return true
		}
		if validarVisitField(visit, context) {
			return true
		}
	}

	return false
}

//validarVisiEdit
func validarVisitField(visit dto.VisitDTO, context *gin.Context) bool {
	context.ShouldBind(&visit)
	msg := utilities.MessageRequired{}
	if len(visit.State) == 0 || visit.State == "" {
		validadRequiredMsg(msg.RequiredState(), context)
		return true
	}
	if len(visit.Description) == 0 || visit.Description == "" {
		validadRequiredMsg(msg.RequiredDescription(), context)
		return true
	}
	if visit.UserId == 0 {
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}
	if visit.SubDetachmentId == 0 {
		validadRequiredMsg(msg.RequiredSubDetachmentId(), context)
		return true
	}
	return false
}
