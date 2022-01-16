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

type ParentService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	All(context *gin.Context)
	AllParentScout(context *gin.Context)
	UserByIdAll(context *gin.Context)
	Remove(context *gin.Context)
}

type parentService struct {
	parentRepository      repositorys.ParentRepository
	userRepository        repositorys.UserRepository
	scoutParentRepository repositorys.ScoutParentRepository
}

func NewParentService() ParentService {
	return &parentService{
		parentRepository:      repositorys.NewParentRepository(),
		userRepository:        repositorys.NewUserRepository(),
		scoutParentRepository: repositorys.NewScoutParentRepository(),
	}
}

//Create
func (service *parentService) Create(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	parent := entity.Parent{}
	parentScout := entity.ParentScout{}
	var parentDto dto.ParentDTO

	findById, _ := service.userRepository.ProfileUser(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.ShouldBind(&parentDto)
	if validarParentCreate(parentDto, context) {
		return
	}
	smapping.FillStruct(&parent, smapping.MapFields(&parentDto))
	existParent, err := service.parentRepository.FindParentByIdentification(parent.Identification)
	if err != nil {
		validadErrors(err, context)
		return
	}
	if existParent.Id > 0 {
		validadExistScout(context)
		return
	}
	data, err := service.parentRepository.CreateParent(parent)
	if err != nil {
		validadErrors(err, context)
		return
	}

	parentScout.ParentId = data.Id
	parentScout.UserId = uint(id)

	_, errs := service.scoutParentRepository.CreateParentScout(parentScout)
	if errs != nil {
		validadErrors(errs, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreateResponse(data))

}

//update
func (service *parentService) Update(context *gin.Context) {
	parent := entity.Parent{}
	var parentDto dto.ParentDTO
	context.ShouldBind(&parentDto)
	if validarParentEdit(parentDto, context) {
		return
	}
	smapping.FillStruct(&parent, smapping.MapFields(&parentDto))
	existId, _ := service.parentRepository.FindParentById(parentDto.Id)
	if existId.Id == 0 {
		validadErrorById(context)
		return
	}
	data, err := service.parentRepository.CreateParent(parent)
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildUpdateResponse(data))
}

func (service *parentService) All(context *gin.Context) {
	var parents, err = service.parentRepository.AllParent()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", parents))
}
func (service *parentService) AllParentScout(context *gin.Context) {
	id, errId := strconv.ParseInt(context.Param("id"), 0, 0)
	if errId != nil {
		validadErrors(errId, context)
		return
	}
	findById, _ := service.userRepository.ProfileUser(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	var parents, err = service.parentRepository.AllParentScout(uint(id))
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", parents))
}

func (service *parentService) UserByIdAll(context *gin.Context) {
	var parents, err = service.parentRepository.AllParent()
	if err != nil {
		validadErrors(err, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(true, "ok", parents))
}

func (service *parentService) Remove(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	findById, _ := service.parentRepository.FindParentById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	parent, err := service.parentRepository.RemoveParent(uint(id))
	if err != nil {
		validadErrorRemove(findById, context)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildDeteleteResponse(parent, findById))
}

//validarRolCreate
func validarParentCreate(r dto.ParentDTO, context *gin.Context) bool {
	context.ShouldBind(&r)
	if len(r.FullName) == 0 || r.FullName == "" {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredName(), context)
		return true
	}
	return false
}

//validarRolCreate
func validarParentEdit(r dto.ParentDTO, context *gin.Context) bool {
	context.ShouldBind(&r)
	msg := utilities.MessageRequired{}
	if r.Id == 0 {
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}
	if len(r.FullName) == 0 || r.FullName == "" {

		validadRequiredMsg(msg.RequiredName(), context)
		return true
	}
	return false
}
