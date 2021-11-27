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

//rolService is a contract.....
type RolService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	FindById(context *gin.Context)
	Delete(context *gin.Context)
	All(context *gin.Context)
	AllRoleModule(context *gin.Context)
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
func (service *rolService) Create(context *gin.Context) {

	rol := entity.Rol{}
	var rolDto dto.RolCreateDTO
	context.ShouldBind(&rolDto)
	if validarRolCreate(rolDto, context) {
		return
	}
	smapping.FillStruct(&rol, smapping.MapFields(&rolDto))

	data, err := service.rolRepository.CreateRol(rol)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (service *rolService) Update(context *gin.Context) {

	rolToUpdate := entity.Rol{}
	var rolDto dto.RolCreateDTO

	context.ShouldBind(&rolDto)
	if validarRolEditar(rolDto, context) {
		return
	}
	err := smapping.FillStruct(&rolToUpdate, smapping.MapFields(&rolDto))
	checkError(err)

	findById, _ := service.rolRepository.FindRolById(uint(rolDto.Id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	data, err := service.rolRepository.UpdateRol(rolToUpdate)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *rolService) All(context *gin.Context) {

	var rols, err = service.rolRepository.AllRol()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", rols)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *rolService) AllRoleModule(context *gin.Context) {

	var roleModule, err = service.rolRepository.RolsModule()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(true, "OK", roleModule)
	context.JSON(http.StatusOK, res)

}
func (service *rolService) Delete(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := service.rolRepository.FindRolById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	status, err := service.rolRepository.DeleteRol(findById)

	if err != nil {
		validadErrorRemove(findById, context)
		return
	}
	res := utilities.BuildDeteleteResponse(status, findById)
	context.JSON(http.StatusOK, res)

}

func (service *rolService) FindById(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}

	rol, err := service.rolRepository.FindRolById(uint(id))
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

//validarRolCreate
func validarRolCreate(r dto.RolCreateDTO, context *gin.Context) bool {
	context.ShouldBind(&r)
	if len(r.Name) == 0 || r.Name == "" {
		msg := utilities.MessageRequired{}
		validadRequiredMsg(msg.RequiredName(), context)
		return true
	}
	return false
}

//validarRolEditar
func validarRolEditar(r dto.RolCreateDTO, context *gin.Context) bool {
	context.ShouldBind(&r)
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
	return false
}
