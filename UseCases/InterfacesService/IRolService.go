package InterfacesService

import "github.com/gin-gonic/gin"

type IRolService interface {
	SetCreateService(context *gin.Context)
	SetUpdateService(context *gin.Context)
	GetFindByIdService(context *gin.Context)
	SetRemoveService(context *gin.Context)
	GetAllService(context *gin.Context)
	GetAllGroupRolService(context *gin.Context)
	GetAllRoleModuleService(context *gin.Context)
}
