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
			log.Println("Claim[id]: ", claims["id"])
			log.Println("Claim[issuer] :", claims["issuer"])
			log.Println("Claim[rol] :", claims["rol"])
		} else {
			log.Println(err)
			response := utilities.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		// authHeader := c.GetHeader("Authorization")

		// if authHeader == "" {
		// 	response := utilities.BuildErrorResponse("Failed to process request", "no token found", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// token, err := jwtService.ValidateToken(authHeader)

		// if token.Valid {

		// 	claims := token.Claims.(jwt.MapClaims)

		// 	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
		// 	fmt.Println(claims["exp"])
		// 	// log.Println("claims[id]", claims["id"])
		// 	// log.Println("claims[issuer]", claims["issuer"])
		// 	// log.Println("claims[exp]", claims["exp"])

		// } else {
		// 	log.Println(err)
		// 	response := utilities.BuildErrorResponse("Token isnot valid", err.Error(), nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// }
	}
}

//AuthorizeJWT validates the token suer given, return 401 if not valid
func GetRol(jwtService services.JWTService, context *gin.Context) (uint, uint) {
	authHeader := context.GetHeader("Authorization")
	token, _ := jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	rol, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["rol"]), 0, 0)

	id, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["id"]), 0, 0)
	return uint(rol), uint(id)
}
