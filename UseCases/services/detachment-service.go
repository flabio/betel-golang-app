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

//DetachmentService is a contract......
type DetachmentService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}

type detachmentService struct {
	detachmentRepository repositorys.DetachmentRepository
}

//NewDetachmentService creates a new instance of RolService
func NewDetachmentService() DetachmentService {
	detachmentRepository := repositorys.NewDetachmentRepository()
	return &detachmentService{
		detachmentRepository: detachmentRepository,
	}
}

func (detachmentService *detachmentService) Create(context *gin.Context) {
	detachment := entity.Detachment{}
	var detachmentDTO dto.DetachmentDTO
	errDTO := context.ShouldBind(&detachmentDTO)
	if errDTO != nil {
		validadErrors(errDTO, context)
		return
	}

	err := smapping.FillStruct(&detachment, smapping.MapFields(&detachmentDTO))
	checkError(err)

	filename, err := UploadFile(context)

	detachment.Archives = filename

	data, err := detachmentService.detachmentRepository.CreateDetachment(detachment)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))

}
func (detachmentService *detachmentService) Update(context *gin.Context) {
	detachment := entity.Detachment{}
	var detachmentDTO dto.DetachmentUpdateDTO

	errDTO := context.ShouldBind(&detachmentDTO)
	if errDTO != nil {
		validadErrors(errDTO, context)
		return
	}
	err := smapping.FillStruct(&detachment, smapping.MapFields(&detachmentDTO))
	checkError(err)

	findId, err := detachmentService.detachmentRepository.FindDetachmentById(detachmentDTO.Id)

	if err != nil {
		validadErrors(err, context)
		return
	}
	if findId.Id == 0 {
		validadErrorById(context)
		return
	}
	filename, err := UploadFile(context)

	if len(findId.Archives) == 0 {

		detachment.Archives = filename
	} else {
		if filename != "" {
			detachment.Archives = filename
		} else {
			detachment.Archives = findId.Archives
		}
	}
	data, err := detachmentService.detachmentRepository.UpdateDetachment(detachment)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(data))
}

//Service of delete detachment
func (detachmentService *detachmentService) Delete(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrorById(context)
		return
	}
	findId, err := detachmentService.detachmentRepository.FindDetachmentById(uint(id))

	if findId.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}

	detachment, err := detachmentService.detachmentRepository.DeleteDetachment(findId)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCanNotDeteleteResponse(detachment)
	context.JSON(http.StatusOK, res)
}

//search by Id derachment
func (detachmentService *detachmentService) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}

	detachment, err := detachmentService.detachmentRepository.FindDetachmentById(uint(id))
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", detachment)
	context.JSON(http.StatusOK, res)
}

//all od detachment
func (detachmentService *detachmentService) All(context *gin.Context) {
	detachment, err := detachmentService.detachmentRepository.AllDetachment()
	if err != nil {
		validadErrors(err, context)
		return
	}
	results := utilities.BuildResponse(true, "OK!", detachment)
	context.JSON(http.StatusOK, results)
}
