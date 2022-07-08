package InterfacesService

import "github.com/gin-gonic/gin"

type IAttendanceService interface {
	CreateAttendanceService(sub_detachment_id uint, context *gin.Context)
	UpdateAttendanceService(context *gin.Context)
	RemoveAttendanceService(context *gin.Context)
	AllAttendanceService(context *gin.Context)
	AttendancesSubdetachmentService(sub_detachment_id uint, context *gin.Context)
	//WeeksbySubDetachmentsAttendanceService(sub_detachment_id uint, context *gin.Context)
}
