package utilities

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptObj struct{}

//BuildResponse method is to inject dasta value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildNotFoudResponse() Response {
	res := Response{
		Status:  false,
		Message: "not found",
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildDanedResponse() Response {
	res := Response{
		Status:  false,
		Message: "permission denied",
		Error:   nil,
	}

	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildExistResponse() Response {
	res := Response{
		Status:  false,
		Message: "Scout already exists.",
		Error:   nil,
	}

	return res
}

//BuildErrorAllResponse method is to inject dasta value to dynamic failed response
func BuildErrorAllResponse(err string) Response {
	res := Response{
		Status:  false,
		Message: err,
	}
	return res
}

//BuildErrorAllResponse method is to inject dasta value to dynamic failed response
func BuildErrorAllResponseMessage(message string) Response {

	res := Response{
		Status:  false,
		Message: message,
	}
	return res
}

//BuildCreateResponse method is to inject dasta value to dynamic success response
func BuildCreateResponse(data interface{}) Response {
	res := Response{
		Status:  true,
		Message: "Create successfully",
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildCreateResponse method is to inject dasta value to dynamic success response
func BuildCreateResp() Response {
	res := Response{
		Status:  true,
		Message: "Create successfully",
		Error:   nil,
	}
	return res
}

//BuildUpdateResponse method is to inject dasta value to dynamic success response
func BuildUpdateResponse(data interface{}) Response {
	res := Response{
		Status:  true,
		Message: "Update successfully",
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildUpdatePasswordResponse method is to inject dasta value to dynamic success response
func BuildUpdatePasswordResponse() Response {
	res := Response{
		Status:  true,
		Message: "Update password successfully",
		Error:   nil,
	}
	return res
}

//BuildUpdatePasswordResponse method is to inject dasta value to dynamic success response
func BuildEmailPasswordIncorrectResponse() Response {
	res := Response{
		Status:  false,
		Message: "Email or Password incorrect",
		Error:   nil,
	}
	return res
}

//BuildUpdateResponse1 method is to inject dasta value to dynamic success response
func BuildUpdateResponses(err string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: "Update successfully",
		Error:   err,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildDeteleteResponse(status bool, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: "It was successfully removed",
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildCanNotDeteleteResponse(data interface{}) Response {

	res := Response{
		Status:  false,
		Message: "The record cannot be deleted",
		Data:    data,
	}
	return res
}

//"Data not found", "No data with given id",
//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildErrorByIdResponse() Response {

	res := Response{
		Status:  false,
		Message: "Data not found",
		Error:   "No data with given id"}
	return res
}

func Pagination(c *gin.Context, limit int) (int, int) {
	p := c.Query("page")

	if p == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		return 1, 0
	}

	begin := (limit * page) - limit
	return page, begin
}
