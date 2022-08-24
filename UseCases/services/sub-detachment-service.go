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

// subDetachmentService
type subDetachmentService struct {
	iSubDetachment Interfaces.ISubDetachment
}

// NewSubDetachmentService
func NewSubDetachmentService() InterfacesService.ISubDetachmentService {
	return &subDetachmentService{
		iSubDetachment: repositorys.NewSubDetachmentRepository(),
	}
}

// Create
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

	res, err := subDetachmentService.iSubDetachment.SetCreateSubDetachment(subDetachment)
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreatedResponse(res))
}

// Update
func (subDetachmentService *subDetachmentService) SetUpdateSubDetachmentService(context *gin.Context) {
	subDetachment := entity.SubDetachment{}
	var dto dto.SubDetachmentDTO
	context.ShouldBind(&dto)
	if validateSubDetachments(dto, context, constantvariables.OPTION_EDIT) {
		return
	}
	smapping.FillStruct(&subDetachment, smapping.MapFields(&dto))

	findById, _ := subDetachmentService.iSubDetachment.GetFindByIdSubDetachment(uint(dto.Id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	filename, err := UploadFile(context)
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if len(findById.Archives) == 0 {
		subDetachment.Archives = filename
	} else {
		if filename != "" {
			subDetachment.Archives = filename
		} else {
			subDetachment.Archives = findById.Archives
		}
	}

	res, err := subDetachmentService.iSubDetachment.SetCreateSubDetachment(subDetachment)
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(res))
}

// Remove
func (subDetachmentService *subDetachmentService) SetRemoveSubDetachmentService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if id == 0 {
		res := utilities.BuildErrResponse(constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	findById, _ := subDetachmentService.iSubDetachment.GetFindByIdSubDetachment(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res, err := subDetachmentService.iSubDetachment.SetRemoveSubDetachment(findById.Id)
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildRemovedResponse(res))

}

// FindById
func (subDetachmentService *subDetachmentService) GetFindByIdSubDetachmentService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if id == 0 {
		res := utilities.BuildErrResponse(constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	findById, _ := subDetachmentService.iSubDetachment.GetFindByIdSubDetachment(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(findById))
}

// FindByIdDetachment
func (subDetachmentService *subDetachmentService) GetFindByIdDetachmentSubDetachmentService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if id == 0 {
		res := utilities.BuildErrResponse(constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res, err := subDetachmentService.iSubDetachment.GetFindByIdDetachmentSubDetachment(uint(id))

	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(res))
}

// All
func (subDetachmentService *subDetachmentService) GetAllSubDetachmentService(context *gin.Context) {
	res, err := subDetachmentService.iSubDetachment.GetAllSubDetachment()
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(res))
}

// Validate
func validateSubDetachments(dto dto.SubDetachmentDTO, context *gin.Context, options int) bool {

	context.ShouldBind(&dto)
	switch options {
	case 1:
		if len(dto.Name) == 0 {
			res := utilities.BuildErrResponse(constantvariables.NAME)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return true
		}
		if dto.DetachmentId == 0 {
			res := utilities.BuildErrResponse(constantvariables.DETACHMENT_ID)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
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
			res := utilities.BuildErrResponse(constantvariables.ID)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return true
		}
		if dto.SubDetachmentId == 0 {
			res := utilities.BuildErrResponse(constantvariables.SUDDETACHMENT)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return true
		}
	}
	return false
}
