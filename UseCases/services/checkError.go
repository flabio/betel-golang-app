package services

import (
	"bete/UseCases/utilities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkError(err error) bool {
	if err != nil {
		log.Fatalf("Failed map %v", err)
		return false

	}
	return true
}

//func of validation
func validadErrors(errDTO error, context *gin.Context) {
	res := utilities.BuildErrorAllResponse(errDTO.Error())
	context.AbortWithStatusJSON(http.StatusBadRequest, res)
}

//validadRequiredMsg
func validadRequiredMsg(message string, context *gin.Context) {

	res := utilities.BuildErrorAllResponseMessage(message)
	context.AbortWithStatusJSON(http.StatusBadRequest, res)
}
func validadErrorById(context *gin.Context) {
	res := utilities.BuildErrorByIdResponse()
	context.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func validadErrorRemove(data interface{}, context *gin.Context) {
	response := utilities.BuildCanNotDeteleteResponse(data)
	context.JSON(http.StatusBadRequest, response)
}
