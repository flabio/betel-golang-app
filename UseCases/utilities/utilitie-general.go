package utilities

import (
	"log"

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
