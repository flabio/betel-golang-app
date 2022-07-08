package InterfacesService

import "github.com/gin-gonic/gin"

type IPatrolService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}
