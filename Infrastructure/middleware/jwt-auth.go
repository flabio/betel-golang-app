package middleware

import (
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/utilities"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GetVariableSession interface {
}
type VariableSession struct {
	Rol      uint
	Id       uint
	Churchid uint
}

var IdUsuario uint

//AuthorizeJWT validates the token suer given, return 401 if not valid
func AuthorizeJWT(jwtService InterfacesService.IJWTService) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.TOKEN_REQUIRED)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) < 2 {
			response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.TOKEN_INVALIDO)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		var extrToken = splitToken[1]

		if len(splitToken) != 2 {
			response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.TOKEN_INVALIDO)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(extrToken)

		if token != nil {
			//token.Claims.(jwt.MapClaims)
			claims := token.Claims.(jwt.MapClaims)
			// log.Println("Claim: ", claims)
			// log.Println("Claim[id]: ", claims["id"])
			// log.Println("Claim[issuer] :", claims["issuer"])
			// log.Println("Claim[rol] :", claims["rol"])
			// log.Println("claims[exp]", claims["exp"])

			//rol, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["rol"]), 0, 0)

			id, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["id"]), 0, 0)
			//churchid, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["churchid"]), 0, 0)
			//IdUsuario=
			u := VariableSession{}
			u.Id = uint(id)
		} else {
			if err != nil {
				response := utilities.BuildErrResponse(http.StatusUnauthorized, constantvariables.TOKEN_NOT_IS_VALIDO)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

		}

	}
}

//AuthorizeJWT validates the token suer given, return 401 if not valid
func ValidadToken(jwtService InterfacesService.IJWTService, context *gin.Context) VariableSession {
	authHeader := context.GetHeader("Authorization")

	if authHeader == "" {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.TOKEN_REQUIRED)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return VariableSession{}
	}
	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 {
		response := utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.TOKEN_INVALIDO)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return VariableSession{}
	}

	var extrToken = splitToken[1]

	token, _ := jwtService.ValidateToken(extrToken)
	claims := token.Claims.(jwt.MapClaims)
	rol, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["rol"]), 0, 0)

	id, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["id"]), 0, 0)
	churchid, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["churchid"]), 0, 0)
	u := VariableSession{
		Id:       uint(id),
		Rol:      uint(rol),
		Churchid: uint(churchid),
	}

	return u
}
