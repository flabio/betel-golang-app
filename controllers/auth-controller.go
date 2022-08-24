package controllers

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/dto"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController intreface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	//Register(ctx *gin.Context)
}

type authController struct {
	authService InterfacesService.IAuthService
	jwtService  InterfacesService.IJWTService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController() AuthController {
	jwtService := services.NewJWTService()
	authService := services.NewAuthService()
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {

	var loginDTO dto.LoginDTO

	ctx.ShouldBind(&loginDTO)

	// if errDTO != nil {
	// 	response := utilities.BuildErrorResponse("Failed to process request", errDTO.Error(), utilities.EmptObj{})
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	// 	return
	// }
	if len(loginDTO.Email) == 0 || len(loginDTO.Password) == 0 {
		response := utilities.BuildErrResponse(constantvariables.PASSWORD_EMAIL_INCORRECT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {

		if v.Roles == nil {
			response := utilities.BuildErrResponse(constantvariables.PARMISSION_DENIED)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.Id), 10), strconv.FormatUint(uint64(v.Roles.RolId), 10), strconv.FormatUint(uint64(v.ChurchId), 10))
		v.Token = generatedToken
		ctx.JSON(http.StatusOK, utilities.BuildResponse(v))
		return
	}

	response := utilities.BuildErrResponse(constantvariables.AGAIN_CREDENTIAL)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// func (c *authController) Register(ctx *gin.Context) {
// 	var registerDTO dto.UserDTO
// 	errDTO := ctx.ShouldBind(&registerDTO)

// 	if errDTO != nil {
// 		response := utilities.BuildErrorResponse("Failed to process request", errDTO.Error(), utilities.EmptObj{})
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	res, err := c.authService.IsDuplicateEmail(registerDTO.Email)
// 	if err != nil {
// 		response := utilities.BuildErrorResponse("Failed to process request", err.Error(), utilities.EmptObj{})
// 		ctx.JSON(http.StatusConflict, response)
// 		return
// 	}
// 	if res {
// 		response := utilities.BuildErrorResponse("Failed to process request", "Duplicate email", utilities.EmptObj{})
// 		ctx.JSON(http.StatusConflict, response)
// 		return
// 	} else {
// 		createdUser := c.authService.CreateUser(registerDTO)
// 		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.Id), 10))
// 		createdUser.Token = token
// 		response := utilities.BuildResponse(true, "OK!", createdUser)
// 		ctx.JSON(http.StatusCreated, response)
// 		return
// 	}
// }
