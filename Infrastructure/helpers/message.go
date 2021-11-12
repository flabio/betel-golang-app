package helpers

import (
	"errors"
)

//Response is used for static shape json return
type Message struct {
	Data interface{} `json:"data"`
}

var (
	errorInt = errors.New("This is int")
)

//BuildResponse method is to inject dasta value to dynamic success response
func getMessage() string {

	return "error dddds"
}
func getString() string {
	return "this is int"
}
func getInt(value int) string {
	return "this is int"
}
func getFloat(value float32) string {
	return "this is float"
}
