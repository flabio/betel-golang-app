package InterfacesService

import "github.com/gin-gonic/gin"

//UserService is a contract.....
type IScoutService interface {
	Create(ChurchId uint, context *gin.Context)
	Update(ChurchId uint, context *gin.Context)
	ListKingsScouts(ChurchId uint, context *gin.Context)
}
