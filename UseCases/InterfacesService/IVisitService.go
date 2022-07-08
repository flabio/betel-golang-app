package InterfacesService

import "github.com/gin-gonic/gin"

type IVisitService interface {
	SetCreateVisitService(context *gin.Context)
	SetUpdateVisitService(context *gin.Context)
	GetAllVisitService(context *gin.Context)
	GetAllVisitByUserVisitService(subDetachmentId uint, context *gin.Context)
	SetRemoveVisitService(context *gin.Context)
}
