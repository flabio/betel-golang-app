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

//Create
func (parentService *parentService) Create(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	parent := entity.Parent{}
	parentScout := entity.ParentScout{}
	var parentDto dto.ParentDTO

	findById, _ := parentService.IUser.GetProfileUser(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.ShouldBind(&parentDto)
	if validarParentCreate(parentDto, context) {
		return
	}
	smapping.FillStruct(&parent, smapping.MapFields(&parentDto))
	existParent, err := parentService.IParent.GetFindParentByIdentification(parent.Identification)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if existParent.Id > 0 {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.SCOUT_EXIST))
		return
	}
	data, err := parentService.IParent.SetCreateParent(parent)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	parentScout.ParentId = data.Id
	parentScout.UserId = uint(id)

	_, errs := parentService.IScoutParent.SetCreateParentScouts(parentScout)
	if errs != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errs.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, data))

}

//update
func (parentService *parentService) Update(context *gin.Context) {
	parent := entity.Parent{}
	var parentDto dto.ParentDTO
	context.ShouldBind(&parentDto)
	if validarParentEdit(parentDto, context) {
		return
	}
	smapping.FillStruct(&parent, smapping.MapFields(&parentDto))
	existId, _ := parentService.IParent.GetFindParentById(parentDto.Id)
	if existId.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := parentService.IParent.SetCreateParent(parent)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, data))
}

func (parentService *parentService) All(context *gin.Context) {
	var parents, err = parentService.IParent.GetAllParent()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "ok", parents))
}
func (parentService *parentService) AllParentScout(context *gin.Context) {
	id, errId := strconv.ParseInt(context.Param("id"), 0, 0)
	if errId != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	findById, _ := parentService.IUser.GetProfileUser(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var parents, err = parentService.IParent.GetAllParentScout(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "ok", parents))
}

func (parentService *parentService) UserByIdAll(context *gin.Context) {
	var parents, err = parentService.IParent.GetAllParent()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, "ok", parents))
}

func (parentService *parentService) Remove(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	findById, _ := parentService.IParent.GetFindParentById(uint(id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	parent, err := parentService.IParent.SetRemoveParent(uint(id))
	if err != nil {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NOT_DELETED)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	if parent {
		context.JSON(http.StatusOK, utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, findById))

	}
}

//validarRolCreate
func validarParentCreate(r dto.ParentDTO, context *gin.Context) bool {
	context.ShouldBind(&r)
	if len(r.FullName) == 0 || r.FullName == "" {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NAME)
		context.JSON(http.StatusBadRequest, response)
		return true
	}
	return false
}

//validarRolCreate
func validarParentEdit(r dto.ParentDTO, context *gin.Context) bool {
	context.ShouldBind(&r)
	if r.Id == 0 {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.JSON(http.StatusBadRequest, response)
		return true
	}
	if len(r.FullName) == 0 || r.FullName == "" {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NAME)
		context.JSON(http.StatusBadRequest, response)
		return true
	}
	return false
}
