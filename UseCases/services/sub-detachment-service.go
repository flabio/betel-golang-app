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

	var dto dto.SubDetachmentRequest
	subDetachment, msg := getMappingSubDetachment(dto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}

	filename, err := UploadFile(context)
	subDetachment.Archives = filename

	res, err := subDetachmentService.iSubDetachment.SetCreateSubDetachment(subDetachment)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreatedResponse(res))
}

// Update
func (subDetachmentService *subDetachmentService) SetUpdateSubDetachmentService(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	var dto dto.SubDetachmentRequest
	subDetachment, msg := getMappingSubDetachment(dto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	findById, _ := subDetachmentService.iSubDetachment.GetFindByIdSubDetachment(uint(id))
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
		subDetachment.Archives = filename
	} else {
		if filename != "" {
			subDetachment.Archives = filename
		} else {
			subDetachment.Archives = findById.Archives
		}
	}

	res, err := subDetachmentService.iSubDetachment.SetUpdateSubDetachment(subDetachment, uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(res))
}

// Remove
func (subDetachmentService *subDetachmentService) SetRemoveSubDetachmentService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}

	findById, _ := subDetachmentService.iSubDetachment.GetFindByIdSubDetachment(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	res, err := subDetachmentService.iSubDetachment.SetRemoveSubDetachment(findById.Id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildRemovedResponse(res))

}

// FindById
func (subDetachmentService *subDetachmentService) GetFindByIdSubDetachmentService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	data, _ := subDetachmentService.iSubDetachment.GetFindByIdSubDetachment(uint(id))
	if data.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(getSubDetachmentResponse(data)))
}

// FindByIdDetachment
func (subDetachmentService *subDetachmentService) GetFindByIdDetachmentSubDetachmentService(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	res, err := subDetachmentService.iSubDetachment.GetFindByIdDetachmentSubDetachment(uint(id))

	if err != nil {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(res))
}

// All
func (subDetachmentService *subDetachmentService) GetAllSubDetachmentService(context *gin.Context) {
	var subDetachmentList []dto.SubDetachmentResponse
	res, err := subDetachmentService.iSubDetachment.GetAllSubDetachment()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, data := range res {
		subDetachmentList = append(subDetachmentList, getSubDetachmentResponse(data))
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(subDetachmentList))
}

// Validate
func getMappingSubDetachment(dto dto.SubDetachmentRequest, context *gin.Context) (entity.SubDetachment, string) {
	subDetachment := entity.SubDetachment{}
	err := context.ShouldBind(&dto)
	if err != nil {
		return subDetachment, utilities.GetMsgErrorRequired(err)
	}
	err = smapping.FillStruct(&subDetachment, smapping.MapFields(&dto))
	if err != nil {
		return subDetachment, err.Error()
	}
	return subDetachment, ""
}

func getSubDetachmentResponse(data entity.SubDetachment) dto.SubDetachmentResponse {
	subDetachment := dto.SubDetachmentResponse{
		Id:             data.Id,
		Name:           data.Name,
		Archives:       data.Archives,
		DetachmentName: data.Detachment.Name,
		DetachmentId:   data.DetachmentId,
		Active:         data.Active,
	}
	return subDetachment
}
