package services

import (
	"log"
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
// func validadErrors(errDTO error, context *gin.Context) {

// 	res := utilities.BuildErrorAllResponse(errDTO.Error())
// 	context.AbortWithStatusJSON(http.StatusBadRequest, res)
// }

// func validadExistScout(context *gin.Context) {

// 	context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildExistResponse())
// }

//validadRequiredMsg
// func validadRequiredMsg(message string, context *gin.Context) {

// 	context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrorAllResponseMessage(message))
// }
// func validadBirdateRequiredMsg(message string, context *gin.Context) {

// 	res := utilities.BuildErrorAllResponseMessage(message)
// 	context.AbortWithStatusJSON(http.StatusBadRequest, res)
// }

// func validadErrorById(context *gin.Context) {
// 	res := utilities.BuildErrorByIdResponse()
// 	context.AbortWithStatusJSON(http.StatusBadRequest, res)
// }

// func validadErrorRemove(data interface{}, context *gin.Context) {
// 	response := utilities.BuildCanNotDeteleteResponse(data)
// 	context.JSON(http.StatusBadRequest, response)
// }
