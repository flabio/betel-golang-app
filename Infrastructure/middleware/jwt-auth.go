package middleware

import (
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GetVariableSession interface {
}
type VariableSession struct {
	Rol             uint
	Id              uint
	Subdetachmentid uint
	Churchid        uint
}

//AuthorizeJWT validates the token suer given, return 401 if not valid
func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utilities.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token != nil {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[id]: ", claims)
			log.Println("Claim[id]: ", claims["id"])
			log.Println("Claim[issuer] :", claims["issuer"])
			log.Println("Claim[rol] :", claims["rol"])
			log.Println("Claim[subdetachmentid] :", claims["subdetachmentid"])
			log.Println("claims[exp]", claims["exp"])
		} else {
			response := utilities.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}

//AuthorizeJWT validates the token suer given, return 401 if not valid
func GetRol(jwtService services.JWTService, context *gin.Context) VariableSession {
	authHeader := context.GetHeader("Authorization")
	token, _ := jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	rol, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["rol"]), 0, 0)

	id, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["id"]), 0, 0)
	subdetachmentid, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["subdetachmentid"]), 0, 0)
	churchid, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["churchid"]), 0, 0)
	u := VariableSession{
		Id:              uint(id),
		Rol:             uint(rol),
		Subdetachmentid: uint(subdetachmentid),
		Churchid:        uint(churchid),
	}

	return u
}
