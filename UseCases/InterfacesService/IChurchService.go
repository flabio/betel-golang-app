package InterfacesService

import "github.com/gin-gonic/gin"

type IChurchService interface {
	CreateChurchService(context *gin.Context)
	UpdateChurch(context *gin.Context)
	DeleteChurch(context *gin.Context)
	FindChurchById(context *gin.Context)
	AllChurch(context *gin.Context)
}
