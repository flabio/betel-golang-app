package InterfacesService

import "github.com/gin-gonic/gin"

type IKanbanService interface {
	GetKanbans(context *gin.Context)
	GetCountKanbans(context *gin.Context)
}
