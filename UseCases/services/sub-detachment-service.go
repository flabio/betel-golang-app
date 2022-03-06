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

//SubDetachmentService
type SubDetachmentService interface {
	SetCreateSubDetachmentService(context *gin.Context)
	SetUpdateSubDetachmentService(context *gin.Context)
	SetRemoveSubDetachmentService(context *gin.Context)
	GetFindByIdSubDetachmentService(context *gin.Context)
	GetFindByIdDetachmentSubDetachmentService(context *gin.Context)
	GetAllSubDetachmentService(context *gin.Context)
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
func (subDetachmentService *subDetachmentService) SetCreateSubDetachmentService(context *gin.Context) {
	subDetachment := entity.SubDetachment{}
	var dto dto.SubDetachmentDTO
	context.ShouldBind(&dto)
	if validateSubDetachments(dto, context, constantvariables.OPTION_CREATE) {
		return
	}
	smapping.FillStruct(&subDetachment, smapping.MapFields(&dto))

	filename, err := UploadFile(context)
	subDetachment.Archives = filename

	res, err := subDetachmentService.subDetachmentRepository.SetCreateSubDetachment(subDetachment)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(res))
}

//Update
func (subDetachmentService *subDetachmentService) SetUpdateSubDetachmentService(context *gin.Context) {
	subDetachment := entity.SubDetachment{}
	var dto dto.SubDetachmentDTO
	context.ShouldBind(&dto)
	if validateSubDetachments(dto, context, constantvariables.OPTION_EDIT) {
		return
	}
	smapping.FillStruct(&subDetachment, smapping.MapFields(&dto))

	findById, _ := subDetachmentService.subDetachmentRepository.GetFindByIdSubDetachment(uint(dto.Id))
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

	res, err := subDetachmentService.subDetachmentRepository.SetCreateSubDetachment(subDetachment)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(res))
}

//Remove
func (subDetachmentService *subDetachmentService) SetRemoveSubDetachmentService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	findById, _ := subDetachmentService.subDetachmentRepository.GetFindByIdSubDetachment(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	res, err := subDetachmentService.subDetachmentRepository.SetRemoveSubDetachment(findById.Id)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(res, findById))

}

//FindById
func (subDetachmentService *subDetachmentService) GetFindByIdSubDetachmentService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	findById, _ := subDetachmentService.subDetachmentRepository.GetFindByIdSubDetachment(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", findById))
}

//FindByIdDetachment
func (subDetachmentService *subDetachmentService) GetFindByIdDetachmentSubDetachmentService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res, err := subDetachmentService.subDetachmentRepository.GetFindByIdDetachmentSubDetachment(uint(id))

	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", res))
}

//All
func (subDetachmentService *subDetachmentService) GetAllSubDetachmentService(context *gin.Context) {
	res, err := subDetachmentService.subDetachmentRepository.GetAllSubDetachment()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", res))
}

//Validate
func validateSubDetachments(dto dto.SubDetachmentDTO, context *gin.Context, options int) bool {

	context.ShouldBind(&dto)
	switch options {
	case 1:
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

	}

	return false
}
func validateUserSubDetachments(dto dto.UserSubDetachmentDTO, context *gin.Context, options int) bool {

	context.ShouldBind(&dto)
	switch options {
	case 1:
		if dto.UserId == 0 {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredId(), context)
			return true
		}
		if dto.SubDetachmentId == 0 {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredDetachment(), context)
			return true
		}
	}
	return false
}
