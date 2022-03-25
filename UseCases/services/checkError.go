package services

import (
	"bete/UseCases/utilities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@param Err is of type error
*/
func checkError(Err error) bool {
	if Err != nil {
		log.Fatalf("Failed map %v", Err)
		return false

	}
	return true
}

//func of validation

/*
@param ErrDTO is of type error
*/
func validadErrors(ErrDTO error, context *gin.Context) {

	res := utilities.BuildErrorAllResponse(ErrDTO.Error())
	context.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func validadExistScout(context *gin.Context) {

	context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildExistResponse())
}

//validadRequiredMsg
/*
@param ErrDTO is of type error
*/
func validadRequiredMsg(Message string, context *gin.Context) {

	context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrorAllResponseMessage(Message))
}
func validadBirdateRequiredMsg(message string, context *gin.Context) {

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
