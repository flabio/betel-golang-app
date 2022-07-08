package InterfacesService

import "github.com/gin-gonic/gin"

type IParentService interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	All(context *gin.Context)
	AllParentScout(context *gin.Context)
	UserByIdAll(context *gin.Context)
	Remove(context *gin.Context)
}
