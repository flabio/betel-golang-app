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

type detachmentService struct {
	IDetachment Interfaces.IDetachment
}

// NewDetachmentService creates a new instance of RolService
func NewDetachmentService() InterfacesService.IDetachmentService {
	return &detachmentService{
		IDetachment: repositorys.GetDetachmentInstance(),
	}
}

func (detachmentService *detachmentService) Create(context *gin.Context) {
	var detachmentDTO dto.DetachmentDTO
	detachment, err := getMappingDetachment(detachmentDTO, context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	filename, err := UploadFile(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	detachment.Archives = filename

	data, err := detachmentService.IDetachment.SetCreateDetachment(detachment)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreatedResponse(data))

}
func (detachmentService *detachmentService) Update(context *gin.Context) {

	var detachmentDTO dto.DetachmentDTO
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	detachment, err := getMappingDetachment(detachmentDTO, context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	findId, err := detachmentService.IDetachment.GetFindDetachmentById(uint(id))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if findId.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	filename, _ := UploadFile(context)

	if len(findId.Archives) == 0 {

		detachment.Archives = filename
	} else {
		if filename != "" {
			detachment.Archives = filename
		} else {
			detachment.Archives = findId.Archives
		}
	}
	data, err := detachmentService.IDetachment.SetUpdateDetachment(detachment, uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(data))
}

// Service of delete detachment
func (detachmentService *detachmentService) Delete(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	findId, err := detachmentService.IDetachment.GetFindDetachmentById(uint(id))

	if findId.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	detachment, err := detachmentService.IDetachment.SetRemoveDetachment(findId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildRemovedResponse(detachment))
}

// search by Id derachment
func (detachmentService *detachmentService) FindById(context *gin.Context) {
	var listDetament []dto.DetachmentListDTO
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	detachment, err := detachmentService.IDetachment.GetFindDetachmentById(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	result := dto.DetachmentListDTO{
		Id:         detachment.Id,
		Name:       detachment.Name,
		Archives:   detachment.Archives,
		Number:     detachment.Number,
		Section:    detachment.Section,
		District:   detachment.District,
		Active:     detachment.Active,
		ChurchName: detachment.Church.Name,
		ChurchId:   detachment.ChurchId,
	}
	listDetament = append(listDetament, result)
	context.JSON(http.StatusOK, utilities.BuildResponse(result))
}

// all od detachment
func (detachmentService *detachmentService) All(context *gin.Context) {
	var listDetament []dto.DetachmentListDTO
	detachment, err := detachmentService.IDetachment.GetAllDetachment()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, data := range detachment {
		result := dto.DetachmentListDTO{
			Id:         data.Id,
			Name:       data.Name,
			Archives:   data.Archives,
			Number:     data.Number,
			Section:    data.Section,
			District:   data.District,
			Active:     data.Active,
			ChurchName: data.Church.Name,
			ChurchId:   data.ChurchId,
		}
		listDetament = append(listDetament, result)
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(listDetament))
}

func getMappingDetachment(detachmentDTO dto.DetachmentDTO, context *gin.Context) (entity.Detachment, error) {
	var detachment entity.Detachment
	err := context.ShouldBind(&detachmentDTO)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return detachment, err
	}

	err = smapping.FillStruct(&detachment, smapping.MapFields(&detachmentDTO))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return detachment, err
	}

	return detachment, nil

}
