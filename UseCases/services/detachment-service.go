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

// Newv creates a new instance of RolService
func NewDetachmentService() InterfacesService.IDetachmentService {
	return &detachmentService{
		IDetachment: repositorys.GetDetachmentInstance(),
	}
}

func (v *detachmentService) Create(context *gin.Context) {
	var detachmentDTO dto.DetachmentRequest
	detachment, msg := getMappingDetachment(detachmentDTO, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	//filename, err := UploadFile(context)
	// if err != nil {
	// 	context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
	// 	return
	// }
	//detachment.Archives = filename
	isExistNumber, _ := v.IDetachment.IsDuplicateNumber(detachment.Number)
	if isExistNumber {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.EXIST_NUMBER))
		return
	}
	data, err := v.IDetachment.SetCreateDetachment(detachment)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreatedResponse(detachmentResponse(data)))
}
func (v *detachmentService) Update(context *gin.Context) {

	var detachmentDTO dto.DetachmentRequest
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	detachment, msg := getMappingDetachment(detachmentDTO, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	findId, err := v.IDetachment.GetFindDetachmentById(uint(id))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if findId.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
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
	data, err := v.IDetachment.SetUpdateDetachment(detachment, uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(detachmentResponse(data)))
}

// Service of delete detachment
func (v *detachmentService) Delete(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	findId, err := v.IDetachment.GetFindDetachmentById(uint(id))

	if findId.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	detachment, err := v.IDetachment.SetRemoveDetachment(findId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildRemovedResponse(detachmentResponse(detachment)))
}

// search by Id derachment
func (v *detachmentService) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	data, err := v.IDetachment.GetFindDetachmentById(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if data.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(detachmentResponse(data)))
}

// all od detachment
func (v *detachmentService) All(context *gin.Context) {
	var listDetament []dto.DetachmentResponse
	detachment, err := v.IDetachment.GetAllDetachment()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	for _, data := range detachment {
		listDetament = append(listDetament, detachmentResponse(data))
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(listDetament))
}

func getMappingDetachment(detachmentDTO dto.DetachmentRequest, context *gin.Context) (entity.Detachment, string) {
	var detachment entity.Detachment
	err := context.ShouldBindJSON(&detachmentDTO)
	if err != nil {
		msgErros := utilities.GetMsgErrorRequired(err)
		return detachment, msgErros
	}
	err = smapping.FillStruct(&detachment, smapping.MapFields(&detachmentDTO))
	if err != nil {
		return detachment, err.Error()
	}
	return detachment, ""
}

func detachmentResponse(data entity.Detachment) dto.DetachmentResponse {

	return dto.DetachmentResponse{
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
}
