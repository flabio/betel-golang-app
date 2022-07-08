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

type rolService struct {
	Irol Interfaces.IRol
}

//27270673- jasinta tello guerrero sds
//NewrolService creates a new instance of rolServicess
func NewRolService() InterfacesService.IRolService {

	return &rolService{
		Irol: repositorys.GetRolInstance(),
	}
}

//service of create
func (rolService *rolService) SetCreateService(context *gin.Context) {

	rolEntity := entity.Rol{}
	var rolDto dto.RolCreateDTO
	context.ShouldBind(&rolDto)
	if validarRol(rolDto, context, constantvariables.OPTION_CREATE) {
		return
	}
	smapping.FillStruct(&rolEntity, smapping.MapFields(&rolDto))

	data, err := rolService.Irol.SetCreateRol(rolEntity)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (rolService *rolService) SetUpdateService(context *gin.Context) {

	rolEntity := entity.Rol{}
	var rolDto dto.RolCreateDTO

	context.ShouldBind(&rolDto)
	if validarRol(rolDto, context, constantvariables.OPTION_EDIT) {
		return
	}
	err := smapping.FillStruct(&rolEntity, smapping.MapFields(&rolDto))
	checkError(err)

	findById, _ := rolService.Irol.GetFindRolById(uint(rolDto.Id))
	if findById.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, err := rolService.Irol.SetCreateRol(rolEntity)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (rolService *rolService) GetAllService(context *gin.Context) {

	var rols, err = rolService.Irol.GetAllRol()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", rols)
	context.JSON(http.StatusOK, res)
}

// service of group AllGroupRol
func (rolService *rolService) GetAllGroupRolService(context *gin.Context) {
	var rols, err = rolService.Irol.GetAllGroupRol()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", rols)
	context.JSON(http.StatusOK, res)
}

//service of all
func (rolService *rolService) GetAllRoleModuleService(context *gin.Context) {

	var roleModule, err = rolService.Irol.GetRolsModule()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", roleModule)
	context.JSON(http.StatusOK, res)

}
func (rolService *rolService) SetRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := rolService.Irol.GetFindRolById(uint(id))
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
	status, err := rolService.Irol.SetRemoveRol(findById)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NOT_DELETED)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if status {
		res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, findById)
		context.JSON(http.StatusOK, res)
	}

}

func (rolService *rolService) GetFindByIdService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	rolById, err := rolService.Irol.GetFindRolById(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if (rolById == entity.Rol{}) {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.GIVEN_ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", rolById)
	context.JSON(http.StatusOK, res)

}

//method private

//validarRol
func validarRol(r dto.RolCreateDTO, context *gin.Context, option int) bool {
	context.ShouldBind(&r)
	switch option {
	case 1:
		if len(r.Name) == 0 || r.Name == "" {
			// msg := utilities.MessageRequired{}
			// validadRequiredMsg(msg.RequiredName(), context)
			res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NAME)
			context.JSON(http.StatusBadRequest, res)
			return true
		}
	case 2:
		if r.Id == 0 {
			// msg := constantvariables._ID
			// validadRequiredMsg(constantvariables._ID, context)
			res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
			context.JSON(http.StatusBadRequest, res)
			return true
		}
		if len(r.Name) == 0 {
			res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NAME)
			context.JSON(http.StatusBadRequest, res)
			return true
		}
	}
	return false
}
