package InterfacesService

import "github.com/gin-gonic/gin"

type IModuleService interface {
	CreateModule(context *gin.Context)
	AddModuleRole(context *gin.Context)
	AllByRoleModule(context *gin.Context)
	UpdateModule(context *gin.Context)
	AllModule(context *gin.Context)
	FindModuleById(context *gin.Context)
	DeleteModule(context *gin.Context)
	DeleteRoleModule(context *gin.Context)
}
