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

// create module
func (moduleService *moduleService) CreateModule(context *gin.Context) {
	var moduledto dto.ModuleDTO
	module, msg := getMappingModule(moduledto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}

	data, err := moduleService.IModule.SetCreateModule(module)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	res := utilities.BuildCreatedResponse(data)
	context.JSON(http.StatusOK, res)
}

// update module
func (moduleService *moduleService) UpdateModule(context *gin.Context) {

	var moduledto dto.ModuleDTO
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	module, msg := getMappingModule(moduledto, context)
	if msg != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(msg))
		return
	}
	//	77506216
	findByModule, _ := moduleService.IModule.GetFindModuleById(uint(id))
	if findByModule.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	res, err := moduleService.IModule.SetUpdateModule(module, uint(id))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, utilities.BuildUpdatedResponse(res))
}

// lists of module
func (moduleService *moduleService) AllModule(context *gin.Context) {
	var modules, err = moduleService.IModule.GetAllModule()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(modules))
}

// find by id module
func (moduleService *moduleService) FindModuleById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	module, err := moduleService.IModule.GetFindModuleById(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	if module.Id == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(module))
}

// delete module
func (moduleService *moduleService) DeleteModule(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	module, err := moduleService.IModule.GetFindModuleById(uint(id))

	if (module == entity.Module{}) {

		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	if err != nil {
		res := utilities.BuildErrResponse(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	status, err := moduleService.IModule.SetRemoveModule(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.NOT_DELETED))
		return
	}
	if status {
		context.JSON(http.StatusOK, utilities.BuildRemovedResponse(module))
	}

}

// add role module
func (moduleService *moduleService) AddModuleRole(context *gin.Context) {
	var moduledto dto.ModuleRoleDTO
	moduleRole, _ := getMappingModuleRole(moduledto, context)
	data, err := moduleService.IModule.SetCreateModuleRole(moduleRole)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildCreatedResponse(data))
}

// AllByRoleModule
func (moduleService *moduleService) AllByRoleModule(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	modules, err := moduleService.IModule.GetAllByRoleModule(uint(id))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, utilities.BuildResponse(modules))
}

// delete rolemodule
func (moduleService *moduleService) DeleteRoleModule(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}
	module, err := moduleService.IModule.GetFindRoleModuleById(uint(id))

	if (module == entity.RoleModule{}) {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.ID))
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrResponse(err.Error()))
		return
	}

	status, err := moduleService.IModule.SetRemoveRoleModule(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.NOT_DELETED))
		return
	}
	if status {
		context.JSON(http.StatusOK, utilities.BuildRemovedResponse(module))
	}
}
func getMappingModuleRole(moduledto dto.ModuleRoleDTO, context *gin.Context) (entity.RoleModule, string) {
	module := entity.RoleModule{}
	err := context.ShouldBind(&moduledto)

	if err != nil {

		msgError := utilities.GetMsgErrorRequired(err)
		return module, msgError
	}
	err = smapping.FillStruct(&module, smapping.MapFields(&moduledto))
	if err != nil {
		return module, err.Error()
	}
	return module, ""
}

func getMappingModule(moduledto dto.ModuleDTO, context *gin.Context) (entity.Module, string) {
	module := entity.Module{}
	err := context.ShouldBind(&moduledto)

	if err != nil {
		msgError := utilities.GetMsgErrorRequired(err)
		return module, msgError
	}
	err = smapping.FillStruct(&module, smapping.MapFields(&moduledto))
	if err != nil {
		return module, err.Error()
	}
	return module, ""
}
