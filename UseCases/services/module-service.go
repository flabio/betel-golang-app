package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type ModuleService interface {
	CreateModule(context *gin.Context)
	AddModuleRole(context *gin.Context)
	UpdateModule(context *gin.Context)
	AllModule(context *gin.Context)
	FindModuleById(context *gin.Context)
	DeleteModule(context *gin.Context)
	DeleteRoleModule(context *gin.Context)
}

type moduleService struct {
	moduleRepository repositorys.ModuleRepository
}

func NewModuleService() ModuleService {
	var moduleRepository = repositorys.NewModuleRepository()
	return &moduleService{
		moduleRepository: moduleRepository,
	}
}
//create module
func (service *moduleService) CreateModule(context *gin.Context) {
	module := entity.Module{}
	var moduledto dto.ModuleDTO
	context.ShouldBind(&moduledto)
	validarModuleCreate(moduledto, context)
	smapping.FillStruct(&module, smapping.MapFields(&moduledto))

	data, err := service.moduleRepository.CreateModule(module)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(data)
	context.JSON(http.StatusOK, res)
}

//update module
func (service *moduleService) UpdateModule(context *gin.Context) {
	module := entity.Module{}
	var moduledto dto.ModuleDTO
	context.ShouldBind(&moduledto)
	fmt.Println(moduledto)
	validarModuleEditar(moduledto, context)

	smapping.FillStruct(&module, smapping.MapFields(&moduledto))
fmt.Println(module)
	res, err := service.moduleRepository.CreateModule(module)

	if err != nil {
		validadErrors(err, context)
		return
	}

	//findByModule, _ := service.moduleRepository.FindModuleById(uint(moduledto.Id))
	//if findByModule.Id == 0 {
	//	validadErrorById(context)
	//	return
	//}
	data := utilities.BuildUpdateResponse(res)
	context.JSON(http.StatusOK, data)

}

//lists of module
func (service *moduleService) AllModule(context *gin.Context) {
	var modules, err = service.moduleRepository.AllModule()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", modules)
	context.JSON(http.StatusOK, res)

}

//find by id module
func (service *moduleService) FindModuleById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	module, err := service.moduleRepository.FindModuleById(uint(id))
	if err != nil {
		validadErrorById(context)
		return
	}
	res := utilities.BuildResponse(true, "OK", module)
	context.JSON(http.StatusOK, res)

}

//delete module
func (service *moduleService) DeleteModule(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	module, err := service.moduleRepository.FindModuleById(uint(id))

	if (module == entity.Module{}) {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}

	status, err := service.moduleRepository.DeleteModule(uint(id))
	if err != nil {
		response := utilities.BuildCanNotDeteleteResponse(module)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	res := utilities.BuildDeteleteResponse(status, module)
	context.JSON(http.StatusOK, res)
}

//add role module
func (service *moduleService) AddModuleRole(context *gin.Context) {
	module := entity.RoleModule{}
	var moduledto dto.ModuleRoleDTO
	errDTO := context.ShouldBind(&moduledto)

	if errDTO != nil {
		validadErrors(errDTO, context)
		return
	}
	err := smapping.FillStruct(&module, smapping.MapFields(&moduledto))
	checkError(err)

	data, err := service.moduleRepository.AddModule(module)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(data)
	context.JSON(http.StatusOK, res)
}

//delete rolemodule
func (service *moduleService) DeleteRoleModule(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		validadErrors(err, context)
		return
	}
	module, err := service.moduleRepository.FindRoleModuleById(uint(id))

	if (module == entity.RoleModule{}) {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}

	status, err := service.moduleRepository.DeleteModule(uint(id))
	if err != nil {
		response := utilities.BuildCanNotDeteleteResponse(module)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	res := utilities.BuildDeteleteResponse(status, module)
	context.JSON(http.StatusOK, res)
}

//validarModuleCreate
func validarModuleCreate(m dto.ModuleDTO, context *gin.Context) bool {
	context.ShouldBind(&m)
	if len(m.Name) == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredName(), context)
		return true
	}
	return false
}

//validarModuleEditar
func validarModuleEditar(m dto.ModuleDTO, context *gin.Context) bool {
	context.ShouldBind(&m)
	if m.Id == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredId(), context)
		return true
	}
	if len(m.Name) == 0 {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredName(), context)
		return true
	}
	return false
}
