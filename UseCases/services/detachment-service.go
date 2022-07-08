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

//NewDetachmentService creates a new instance of RolService
func NewDetachmentService() InterfacesService.IDetachmentService {
	return &detachmentService{
		IDetachment: repositorys.GetDetachmentInstance(),
	}
}

func (detachmentService *detachmentService) Create(context *gin.Context) {
	detachment := entity.Detachment{}
	var detachmentDTO dto.DetachmentDTO
	errDTO := context.ShouldBind(&detachmentDTO)
	if errDTO != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errDTO.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := smapping.FillStruct(&detachment, smapping.MapFields(&detachmentDTO))
	checkError(err)

	filename, err := UploadFile(context)

	detachment.Archives = filename

	data, err := detachmentService.IDetachment.SetCreateDetachment(detachment)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, data))

}
func (detachmentService *detachmentService) Update(context *gin.Context) {
	detachment := entity.Detachment{}
	var detachmentDTO dto.DetachmentUpdateDTO

	errDTO := context.ShouldBind(&detachmentDTO)
	if errDTO != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errDTO.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}
	err := smapping.FillStruct(&detachment, smapping.MapFields(&detachmentDTO))
	checkError(err)

	findId, err := detachmentService.IDetachment.GetFindDetachmentById(detachmentDTO.Id)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if findId.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
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
	data, err := detachmentService.IDetachment.SetCreateDetachment(detachment)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, data))
}

//Service of delete detachment
func (detachmentService *detachmentService) Delete(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	findId, err := detachmentService.IDetachment.GetFindDetachmentById(uint(id))

	if findId.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	detachment, err := detachmentService.IDetachment.SetRemoveDetachment(findId)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, detachment)
	context.JSON(http.StatusOK, res)
}

//search by Id derachment
func (detachmentService *detachmentService) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	detachment, err := detachmentService.IDetachment.GetFindDetachmentById(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", detachment)
	context.JSON(http.StatusOK, res)
}

//all od detachment
func (detachmentService *detachmentService) All(context *gin.Context) {
	detachment, err := detachmentService.IDetachment.GetAllDetachment()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	results := utilities.BuildResponse(http.StatusOK, "OK!", detachment)
	context.JSON(http.StatusOK, results)
}
