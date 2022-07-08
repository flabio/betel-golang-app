package InterfacesService

import "github.com/gin-gonic/gin"

type ICityService interface {
	All(context *gin.Context)
}
