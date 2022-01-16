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

type VisitService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	All(context *gin.Context)
	AllVisitByUser(subDetachmentId uint, context *gin.Context)
	Remove(context *gin.Context)
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
func (service *visitService) Create(context *gin.Context) {
	var visitDto dto.VisitDTO
	var visit entity.Visit
	context.ShouldBind(&visitDto)
	if validarVisiCreate(visitDto, context) {
		return
	}
	smapping.FillStruct(&visit, smapping.MapFields(&visitDto))
	data, err := service.visitRepository.Create(visit)

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))
}

//Update
func (service *visitService) Update(context *gin.Context) {
	var visitDto dto.VisitDTO
	var visit entity.Visit
	context.ShouldBind(&visitDto)
	if validarVisiEdit(visitDto, context) {
		return
	}
	smapping.FillStruct(&visit, smapping.MapFields(&visitDto))
	existVisit, _ := service.visitRepository.FindById(visitDto.Id)
	if existVisit.Id == 0 {
		validadErrorById(context)
		return
	}
	data, err := service.visitRepository.Create(visit)

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))
}

//All visit
func (service *visitService) All(context *gin.Context) {
	data, err := service.visitRepository.All()

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", data))
}

//AllVisitByUser by iduser and idsubdetachment
func (service *visitService) AllVisitByUser(subDetachmentId uint, context *gin.Context) {
	id, errId := strconv.ParseInt(context.Param("id"), 0, 0)
	if errId != nil {
		validadErrors(errId, context)
		return
	}
	data, err := service.visitRepository.AllVisitByUser(uint(id), uint(subDetachmentId))

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", data))

}

func (service *visitService) Remove(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	existVisit, _ := service.visitRepository.FindById(uint(id))
	if existVisit.Id == 0 {
		validadErrorById(context)
		return
	}
	parent, err := service.visitRepository.Remove(uint(id))
	if err != nil {
		validadErrorRemove(id, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(parent, existVisit))
}

//validarVisiCreate
func validarVisiCreate(visit dto.VisitDTO, context *gin.Context) bool {
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

//validarVisiEdit
func validarVisiEdit(visit dto.VisitDTO, context *gin.Context) bool {
	context.ShouldBind(&visit)
	msg := utilities.MessageRequired{}
	if visit.Id == 0 {
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}
	if len(visit.State) == 0 || visit.State == "" {
		validadRequiredMsg(msg.RequiredState(), context)
		return true
	}
	if len(visit.Description) == 0 || visit.Description == "" {
		validadRequiredMsg(msg.RequiredDescription(), context)
		return true
	}
	if visit.UserId == 0 {
		validadRequiredMsg(msg.RequiredUserId(), context)
		return true
	}
	if visit.SubDetachmentId == 0 {
		validadRequiredMsg(msg.RequiredSubDetachmentId(), context)
		return true
	}
	return false
}
