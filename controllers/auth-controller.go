package controllers

import (
	"bete/Core/entity"
	"bete/UseCases/dto"
	"bete/UseCases/services"
	"bete/UseCases/utilities"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AuthController intreface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	//Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

//NewAuthController creates a new instance of AuthController
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
		response := utilities.BuildEmailPasswordIncorrectResponse()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {
		//
		RoldId := c.authService.GetIdRol(v.Id)

		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.Id), 10), strconv.FormatUint(uint64(RoldId), 10), strconv.FormatUint(uint64(v.SubDetachmentId), 10), strconv.FormatUint(uint64(v.ChurchId), 10))
		v.Token = generatedToken
		response := utilities.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := utilities.BuildErrorResponse("Please check again your credential", "Invalid Credential", utilities.EmptObj{})
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
