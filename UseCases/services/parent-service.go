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

type parentService struct {
	IParent      Interfaces.IParent
	IUser        Interfaces.IUser
	IScoutParent Interfaces.IScoutParent
}

func NewParentService() InterfacesService.IParentService {
	return &parentService{
		IParent:      repositorys.NewParentRepository(),
		IUser:        repositorys.NewUserRepository(),
		IScoutParent: repositorys.GetScoutParentInstance(),
	}
}

// Create
func (parentService *parentService) Create(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	parentScout := entity.ParentScout{}
	var parentDto dto.ParentDTO

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	parent, err := getMappingParent(parentDto, context)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	existParent, err := parentService.IParent.GetFindParentByIdentification(parent.Identification)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if existParent.Id > 0 {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.SCOUT_EXIST))
		return
	}
	data, err := parentService.IParent.SetCreateParent(parent)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	parentScout.ParentId = data.Id
	parentScout.UserId = uint(id)

	_, errs := parentService.IScoutParent.SetCreateParentScouts(parentScout)
	if errs != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(errs.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildUpdatedResponse(data))
}

// update
func (parentService *parentService) Update(context *gin.Context) {
	var parentDto dto.ParentDTO
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	parent, err := getMappingParent(parentDto, context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	existId, _ := parentService.IParent.GetFindParentById(uint(id))
	if existId.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	data, err := parentService.IParent.SetCreateParent(parent)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, utilities.BuildUpdatedResponse(data))
}

func (parentService *parentService) All(context *gin.Context) {
	var parents, err = parentService.IParent.GetAllParent()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(parents))
}
func (parentService *parentService) AllParentScout(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	findById, _ := parentService.IUser.GetProfileUser(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}

	parents, err := parentService.IParent.GetAllParentScout(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(parents))
}

func (parentService *parentService) UserByIdAll(context *gin.Context) {
	var parents, err = parentService.IParent.GetAllParent()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(parents))
}

func (parentService *parentService) Remove(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	findById, _ := parentService.IParent.GetFindParentById(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	parent, err := parentService.IParent.SetRemoveParent(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.NOT_DELETED))
		return
	}
	if parent {
		context.JSON(http.StatusCreated, utilities.BuildRemovedResponse(findById))
	}
}

func getMappingParent(dto dto.ParentDTO, context *gin.Context) (entity.Parent, error) {
	var parent entity.Parent
	err := context.ShouldBind(&dto)
	if err != nil {
		return parent, err
	}
	err = smapping.FillStruct(&parent, smapping.MapFields(&dto))
	if err != nil {
		return parent, err
	}

	return parent, nil
}
