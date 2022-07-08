package InterfacesService

import "github.com/gin-gonic/gin"

//UserService is a contract.....
type IUserService interface {
	SetCreateService(context *gin.Context)

	SetUpdateService(context *gin.Context)
	SetUpdatePasswordService(context *gin.Context)
	GetAllService(context *gin.Context)
	GetListUserService(context *gin.Context)
	GetListKingsScoutsService(context *gin.Context)
	SetRemoveService(context *gin.Context)
	GetProfileService(userID uint, context *gin.Context)
	GetFindUserService(context *gin.Context)
	GetFindUserNameLastNameService(context *gin.Context)
}
