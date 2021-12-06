package services

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(context *gin.Context) (string, error) {

	file, errs := context.FormFile("file")
	if errs != nil {

		return "", errs
	}
	if file == nil {
		return "", nil
	}
	filenameNew := filepath.Base(file.Filename)
	now := time.Now()

	filenameNew = strconv.Itoa(now.Nanosecond()) + "" + filenameNew
	errpath := context.SaveUploadedFile(file, "assets/"+filenameNew)

	if errpath != nil {
		return "", errpath
	}
	return "localhost:8080/assets/" + filenameNew, nil
}

func UploadFileDocument(context *gin.Context) (string, error) {

	file, errs := context.FormFile("document_identification")
	if errs != nil {

		return "", errs
	}
	if file == nil {
		return "", nil
	}
	filenameNew := filepath.Base(file.Filename)
	now := time.Now()

	filenameNew = strconv.Itoa(now.Nanosecond()) + "" + filenameNew
	errpath := context.SaveUploadedFile(file, "assets/document/"+filenameNew)

	if errpath != nil {
		return "", errpath
	}
	return "localhost:8080/assets/" + filenameNew, nil
}
