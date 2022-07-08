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

type moduleService struct {
	IModule Interfaces.IModule
}

func NewModuleService() InterfacesService.IModuleService {
	moduleRepository := repositorys.NewModuleRepository()
	return &moduleService{
		IModule: moduleRepository,
	}
}

//create module
func (moduleService *moduleService) CreateModule(context *gin.Context) {
	module := entity.Module{}
	var moduledto dto.ModuleDTO
	context.ShouldBind(&moduledto)
	validarModuleCreate(moduledto, context)
	smapping.FillStruct(&module, smapping.MapFields(&moduledto))

	data, err := moduleService.IModule.SetCreateModule(module)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, data)
	context.JSON(http.StatusOK, res)
}

//update module
func (moduleService *moduleService) UpdateModule(context *gin.Context) {
	module := entity.Module{}
	var moduledto dto.ModuleDTO
	context.ShouldBind(&moduledto)
	validarModuleEditar(moduledto, context)

	smapping.FillStruct(&module, smapping.MapFields(&moduledto))

	res, err := moduleService.IModule.SetCreateModule(module)

	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	findByModule, _ := moduleService.IModule.GetFindModuleById(uint(moduledto.Id))
	if findByModule.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_UPDATE, res)
	context.JSON(http.StatusOK, data)

}

//lists of module
func (moduleService *moduleService) AllModule(context *gin.Context) {
	var modules, err = moduleService.IModule.GetAllModule()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "", modules)
	context.JSON(http.StatusOK, res)

}

//find by id module
func (moduleService *moduleService) FindModuleById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	module, err := moduleService.IModule.GetFindModuleById(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", module)
	context.JSON(http.StatusOK, res)

}

//delete module
func (moduleService *moduleService) DeleteModule(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	module, err := moduleService.IModule.GetFindModuleById(uint(id))

	if (module == entity.Module{}) {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	status, err := moduleService.IModule.SetRemoveModule(uint(id))
	if err != nil {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NOT_DELETED)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	if status {
		res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, module)
		context.JSON(http.StatusOK, res)
	}

}

//add role module
func (moduleService *moduleService) AddModuleRole(context *gin.Context) {
	module := entity.RoleModule{}
	var moduledto dto.ModuleRoleDTO
	errDTO := context.ShouldBind(&moduledto)

	if errDTO != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, errDTO.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err := smapping.FillStruct(&module, smapping.MapFields(&moduledto))
	checkError(err)

	data, err := moduleService.IModule.SetCreateModuleRole(module)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_CREATE, data)
	context.JSON(http.StatusOK, res)
}

//AllByRoleModule
func (moduleService *moduleService) AllByRoleModule(context *gin.Context) {
	id, errid := strconv.ParseUint(context.Param("id"), 0, 0)
	if errid != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var modules, err = moduleService.IModule.GetAllByRoleModule(uint(id))
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utilities.BuildResponse(http.StatusOK, "OK", modules)
	context.JSON(http.StatusOK, res)
}

//delete rolemodule
func (moduleService *moduleService) DeleteRoleModule(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	module, err := moduleService.IModule.GetFindRoleModuleById(uint(id))

	if (module == entity.RoleModule{}) {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	status, err := moduleService.IModule.SetRemoveRoleModule(uint(id))
	if err != nil {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NOT_DELETED)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	if status {

		res := utilities.BuildResponse(http.StatusOK, constantvariables.SUCCESS_IT_WAS_REMOVED, module)
		context.JSON(http.StatusOK, res)
	}
}

//validarModuleCreate
func validarModuleCreate(m dto.ModuleDTO, context *gin.Context) bool {
	context.ShouldBind(&m)
	if len(m.Name) == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	return false
}

//validarModuleEditar
func validarModuleEditar(m dto.ModuleDTO, context *gin.Context) bool {
	context.ShouldBind(&m)
	if m.Id == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.ID)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	if len(m.Name) == 0 {
		res := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.NAME)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return true
	}
	return false
}
