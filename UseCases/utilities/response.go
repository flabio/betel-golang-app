package utilities

import (
	constantvariables "bete/Infrastructure/constantVariables"
)

//Response is used for static shape json return
type Response struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//Response is used for static shape json return
type ResponseErr struct {
	Status  int64  `json:"status"`
	Message string `json:"error"`
	ok      bool   `json:"status"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptObj struct{}

//BuildResponse method is to inject dasta value to dynamic success response
func BuildResponse(status int, message string, data interface{}) Response {
	res := Response{
		Status:  int64(status),
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
// func BuildErrorResponse(message string, err string, data interface{}) Response {
// 	splittedError := strings.Split(err, "\n")
// 	res := Response{
// 		Status:  400,
// 		Message: message,
// 		Error:   splittedError,
// 		Data:    data,
// 	}
// 	return res
// }

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildErrResponse(status int, message string) ResponseErr {

	res := ResponseErr{
		Status:  int64(status),
		Message: message,
		ok:      false,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildNotFoudResponse() Response {
	res := Response{
		Status:  403,
		Message: "not found",
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
// func BuildDanedResponse() Response {
// 	res := Response{
// 		Status:  404,
// 		Message: "permission denied",
// 		Error:   nil,
// 	}

// 	return res
// }

//BuildErrorResponse method is to inject dasta value to dynamic failed response
// func BuildExistResponse() Response {
// 	res := Response{
// 		Status:  false,
// 		Message: "Scout already exists.",
// 		Error:   nil,
// 	}

// 	return res
// }

//BuildErrorAllResponse method is to inject dasta value to dynamic failed response
// func BuildErrorAllResponse(err string) Response {
// 	res := Response{
// 		Status:  false,
// 		Message: err,
// 	}
// 	return res
// }

//BuildErrorAllResponse method is to inject dasta value to dynamic failed response
// func BuildErrorAllResponseMessage(message string) Response {

// 	res := Response{
// 		Status:  false,
// 		Message: message,
// 	}
// 	return res
// }

//BuildCreateResponse method is to inject dasta value to dynamic success response
func BuildCreateResponse(status int, data interface{}) Response {
	res := Response{
		Status:  int64(status),
		Message: constantvariables.SUCCESS_CREATE,
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildCreateResponse method is to inject dasta value to dynamic success response
// func BuildCreateResp() Response {
// 	res := Response{
// 		Status:  true,
// 		Message: constantvariables.SUCCESS_CREATE,
// 		Error:   nil,
// 	}
// 	return res
// }

//BuildUpdateResponse method is to inject dasta value to dynamic success response
// func BuildUpdateResponse(data interface{}) Response {
// 	res := Response{
// 		Status:  true,
// 		Message: constantvariables.SUCCESS_UPDATE,
// 		Error:   nil,
// 		Data:    data,
// 	}
// 	return res
// }

//BuildUpdatePasswordResponse method is to inject dasta value to dynamic success response
// func BuildUpdatePasswordResponse() Response {
// 	res := Response{
// 		Status:  true,
// 		Message: constantvariables.SUCCESS_PASSWORD_UPDATE,
// 		Error:   nil,
// 	}
// 	return res
// }

//BuildUpdatePasswordResponse method is to inject dasta value to dynamic success response
// func BuildEmailPasswordIncorrectResponse() Response {
// 	res := Response{
// 		Status:  false,
// 		Message: constantvariables.PASSWORD_EMAIL_INCORRECT,
// 		Error:   nil,
// 	}
// 	return res
// }

//BuildUpdateResponse1 method is to inject dasta value to dynamic success response
// func BuildUpdateResponses(err string, data interface{}) Response {
// 	res := Response{
// 		Status:  true,
// 		Message: constantvariables.SUCCESS_UPDATE,
// 		Error:   err,
// 		Data:    data,
// 	}
// 	return res
// }

//BuildErrorResponse method is to inject dasta value to dynamic failed response
// func BuildDeteleteResponse(status bool, data interface{}) Response {
// 	res := Response{
// 		Status:  status,
// 		Message: constantvariables.SUCCESS_IT_WAS_REMOVED,
// 		Data:    data,
// 	}
// 	return res
// }

//BuildErrorResponse method is to inject dasta value to dynamic failed response
// func BuildCanNotDeteleteResponse(data interface{}) Response {

// 	res := Response{
// 		Status:  false,
// 		Message: constantvariables.NOT_DELETED,
// 		Data:    data,
// 	}
// 	return res
// }

// //"Data not found", "No data with given id",
// //BuildErrorResponse method is to inject dasta value to dynamic failed response
// func BuildErrorByIdResponse() Response {

// 	res := Response{
// 		Status:  false,
// 		Message: "Data not found",
// 		Error:   "No data with given id"}
// 	return res
// }
