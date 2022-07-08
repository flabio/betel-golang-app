package InterfacesService

import "github.com/gin-gonic/gin"

type IDetachmentService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	All(context *gin.Context)
}
