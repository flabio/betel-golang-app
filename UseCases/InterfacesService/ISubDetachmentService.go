package InterfacesService

import "github.com/gin-gonic/gin"

//SubDetachmentService
type ISubDetachmentService interface {
	SetCreateSubDetachmentService(context *gin.Context)
	SetUpdateSubDetachmentService(context *gin.Context)
	SetRemoveSubDetachmentService(context *gin.Context)
	GetFindByIdSubDetachmentService(context *gin.Context)
	GetFindByIdDetachmentSubDetachmentService(context *gin.Context)
	GetAllSubDetachmentService(context *gin.Context)
}
