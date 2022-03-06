package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/UseCases/dto"
	"bete/UseCases/utilities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

//rolService is a contract.....
type RolService interface {
	SetCreateService(context *gin.Context)
	SetUpdateService(context *gin.Context)
	GetFindByIdService(context *gin.Context)
	SetRemoveService(context *gin.Context)
	GetAllService(context *gin.Context)
	GetAllGroupRolService(context *gin.Context)
	GetAllRoleModuleService(context *gin.Context)
}

type rolService struct {
	rolRepository repositorys.RolRepository
}

//27270673- jasinta tello guerrero sds
//NewrolService creates a new instance of rolServicess
func NewRolService() RolService {

	rolRepository := repositorys.NewRolRepository()
	return &rolService{
		rolRepository: rolRepository,
	}
}

//service of create
func (service *rolService) SetCreateService(context *gin.Context) {

	rol := entity.Rol{}
	var rolDto dto.RolCreateDTO
	context.ShouldBind(&rolDto)
	if validarRol(rolDto, context, constantvariables.OPTION_CREATE) {
		return
	}
	smapping.FillStruct(&rol, smapping.MapFields(&rolDto))

	data, err := service.rolRepository.SetCreateRol(rol)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (service *rolService) SetUpdateService(context *gin.Context) {

	rolToUpdate := entity.Rol{}
	var rolDto dto.RolCreateDTO

	context.ShouldBind(&rolDto)
	if validarRol(rolDto, context, constantvariables.OPTION_EDIT) {
		return
	}
	err := smapping.FillStruct(&rolToUpdate, smapping.MapFields(&rolDto))
	checkError(err)

	findById, _ := service.rolRepository.GetFindRolById(uint(rolDto.Id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	data, err := service.rolRepository.SetCreateRol(rolToUpdate)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *rolService) GetAllService(context *gin.Context) {

	var rols, err = service.rolRepository.GetAllRol()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", rols)
	context.JSON(http.StatusOK, res)
}

// service of group AllGroupRol
func (service *rolService) GetAllGroupRolService(context *gin.Context) {
	var rols, err = service.rolRepository.GetAllGroupRol()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", rols)
	context.JSON(http.StatusOK, res)
}

//service of all
func (service *rolService) GetAllRoleModuleService(context *gin.Context) {

	var roleModule, err = service.rolRepository.GetRolsModule()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", roleModule)
	context.JSON(http.StatusOK, res)

}
func (service *rolService) SetRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := service.rolRepository.GetFindRolById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	status, err := service.rolRepository.SetRemoveRol(findById)

	if err != nil {
		validadErrorRemove(findById, context)
		return
	}
	res := utilities.BuildDeteleteResponse(status, findById)
	context.JSON(http.StatusOK, res)

}

func (service *rolService) GetFindByIdService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	fmt.Println("aqui servicios")
	rol, err := service.rolRepository.GetFindRolById(uint(id))
	if err != nil {
		validadErrors(err, context)
		return
	}

	if (rol == entity.Rol{}) {
		validadErrorById(context)
		return
	}
	res := utilities.BuildResponse(true, "OK", rol)
	context.JSON(http.StatusOK, res)

}

//method private

//validarRol
func validarRol(r dto.RolCreateDTO, context *gin.Context, option int) bool {
	context.ShouldBind(&r)
	switch option {
	case 1:
		if len(r.Name) == 0 || r.Name == "" {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredName(), context)
			return true
		}
	case 2:
		if r.Id == 0 {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredId(), context)
			return true
		}
		if len(r.Name) == 0 {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredName(), context)
			return true
		}
	}
	return false
}
