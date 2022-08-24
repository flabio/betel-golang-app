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

// 27270673- jasinta tello guerrero sds
// NewrolService creates a new instance of rolServicess
func NewRolService() InterfacesService.IRolService {

	return &rolService{
		Irol: repositorys.GetRolInstance(),
	}
}

// service of create
func (rolService *rolService) SetCreateService(context *gin.Context) {

	var rolDto dto.RolDTO
	rol, err := getMappingRol(rolDto, context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	data, err := rolService.Irol.SetCreateRol(rol)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	res := utilities.BuildCreatedResponse(data)
	context.JSON(http.StatusCreated, res)
}

// service of update
func (rolService *rolService) SetUpdateService(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	var rolDto dto.RolDTO

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	rol, err := getMappingRol(rolDto, context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	findById, _ := rolService.Irol.GetFindRolById(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}

	data, err := rolService.Irol.SetCreateRol(rol)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	res := utilities.BuildUpdatedResponse(data)
	context.JSON(http.StatusCreated, res)

}

// service of all
func (rolService *rolService) GetAllService(context *gin.Context) {

	var rols, err = rolService.Irol.GetAllRol()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	res := utilities.BuildResponse(rols)
	context.JSON(http.StatusOK, res)
}

// service of group AllGroupRol
func (rolService *rolService) GetAllGroupRolService(context *gin.Context) {
	var rols, err = rolService.Irol.GetAllGroupRol()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	res := utilities.BuildResponse(rols)
	context.JSON(http.StatusOK, res)
}

// service of all
func (rolService *rolService) GetAllRoleModuleService(context *gin.Context) {

	var roleModule, err = rolService.Irol.GetRolsModule()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	res := utilities.BuildResponse(roleModule)
	context.JSON(http.StatusOK, res)

}
func (rolService *rolService) SetRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	findById, _ := rolService.Irol.GetFindRolById(uint(id))
	if findById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	status, err := rolService.Irol.SetRemoveRol(findById)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.NOT_DELETED))
		return
	}
	if status {
		context.JSON(http.StatusCreated, utilities.BuildRemovedResponse(findById))
	}

}

func (rolService *rolService) GetFindByIdService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	rolById, err := rolService.Irol.GetFindRolById(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	if rolById.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.GIVEN_ID))
		return
	}
	res := utilities.BuildResponse(rolById)
	context.JSON(http.StatusOK, res)

}

//method private

func getMappingRol(rolDto dto.RolDTO, context *gin.Context) (entity.Rol, error) {
	rol := entity.Rol{}
	err := context.ShouldBind(&rolDto)

	if err != nil {
		return rol, err
	}

	err = smapping.FillStruct(&rol, smapping.MapFields(&rolDto))
	if err != nil {
		return rol, err
	}
	return rol, nil

}
