package controllers

import (
	constantvariables "bete/Infrastructure/constantVariables"
	"bete/Infrastructure/middleware"

	"bete/UseCases/InterfacesService"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// scoutController is a ....
type ScoutController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	//Delete(context *gin.Context)
	ListKingsScouts(context *gin.Context)
}

type scoutController struct {
	user InterfacesService.IScoutService
	jwt  InterfacesService.IJWTService
}

// NewscoutController is creating anew instance of UserControlller
func NewScoutController() ScoutController {
	return &scoutController{
		user: services.NewScoutService(),
		jwt:  services.NewJWTService(),
	}
}

// create user method post
func (c *scoutController) Create(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.user.Create(claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// update user method push
func (c *scoutController) Update(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.user.Update(claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))

}

func (c *scoutController) ListKingsScouts(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.user.ListKingsScouts(claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

// //delete user
// func (c *scoutController) Delete(context *gin.Context) {

// 	rol, _ := middleware.ValidadToken(c.jwt, context)
// 	if rol > 0 {
// 		c.user.Delete(context)
// 		return
// 	}
// 	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse( constantvariables.PERMISSION_DANIED))

// }
