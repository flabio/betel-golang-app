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

//SubDetachmentService
type SubDetachmentService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}

//subDetachmentService
type subDetachmentService struct {
	subDetachmentRepository repositorys.SubDetachmentRepository
}

//NewSubDetachmentService
func NewSubDetachmentService() SubDetachmentService {
	subDetachmentRepository := repositorys.NewSubDetachmentRepository()
	return &subDetachmentService{
		subDetachmentRepository: subDetachmentRepository,
	}
}

//Create
func (subDetachmentService *subDetachmentService) Create(context *gin.Context) {
	subDetachment := entity.SubDetachment{}
	var dto dto.SubDetachmentDTO
	context.ShouldBind(&dto)
	if validateSubDetachments(dto, context) {
		return
	}
	smapping.FillStruct(&subDetachment, smapping.MapFields(&dto))

	filename, err := UploadFile(context)
	subDetachment.Archives = filename

	res, err := subDetachmentService.subDetachmentRepository.Create(subDetachment)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(res))
}

//Update
func (subDetachmentService *subDetachmentService) Update(context *gin.Context) {
	subDetachment := entity.SubDetachment{}
	var dto dto.SubDetachmentDTO
	context.ShouldBind(&dto)
	if validateSubDetachments(dto, context) {
		return
	}
	smapping.FillStruct(&subDetachment, smapping.MapFields(&dto))

	findById, _ := subDetachmentService.subDetachmentRepository.FindById(uint(dto.Id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	filename, err := UploadFile(context)

	if len(findById.Archives) == 0 {
		subDetachment.Archives = filename
	} else {
		if filename != "" {
			subDetachment.Archives = filename
		} else {
			subDetachment.Archives = findById.Archives
		}
	}

	res, err := subDetachmentService.subDetachmentRepository.Update(subDetachment)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(res))
}

//Remove
func (subDetachmentService *subDetachmentService) Remove(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	findById, _ := subDetachmentService.subDetachmentRepository.FindById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	res, err := subDetachmentService.subDetachmentRepository.Remove(findById.Id)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(res, findById))

}

//FindById
func (subDetachmentService *subDetachmentService) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	findById, _ := subDetachmentService.subDetachmentRepository.FindById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", findById))
}

//All
func (subDetachmentService *subDetachmentService) All(context *gin.Context) {
	res, err := subDetachmentService.subDetachmentRepository.All()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", res))
}

//Validate
func validateSubDetachments(dto dto.SubDetachmentDTO, context *gin.Context) bool {

	context.ShouldBind(&dto)
	if len(dto.Name) == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}
	if dto.DetachmentId == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredDetachment(), context)
		return true
	}
	return false
}
