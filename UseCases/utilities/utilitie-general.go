package utilities

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		checkError(err)
	}
	return string(hash)
}
func checkError(err error) bool {
	if err != nil {
		log.Fatalf("Failed map %v", err)
		return false

	}
	return true
}

func GetMsgErrorRequired(err error) string {
	var verr validator.ValidationErrors
	e := errors.As(err, &verr)
	if e {
		for _, f := range verr {
			return f.Field() + " is " + f.Tag()
		}
	}
	return err.Error()
}
