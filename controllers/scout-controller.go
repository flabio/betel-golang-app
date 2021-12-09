package controllers

import (
	"bete/Infrastructure/middleware"

	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

//scoutController is a ....
type ScoutController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	//Delete(context *gin.Context)
	ListKingsScouts(context *gin.Context)
}

type scoutController struct {
	user services.ScoutService
	role services.UserRolService
	jwt  services.JWTService
}

//NewscoutController is creating anew instance of UserControlller
func NewScoutController() ScoutController {
	return &scoutController{
		user: services.NewScoutService(),
		role: services.NewUserRolService(),
		jwt:  services.NewJWTService(),
	}
}

//create user method post
func (c *scoutController) Create(context *gin.Context) {
	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.user.Create(claim.Subdetachmentid, claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//update user method push
func (c *scoutController) Update(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.user.Update(claim.Subdetachmentid, claim.Churchid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

func (c *scoutController) ListKingsScouts(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.user.ListKingsScouts(claim.Subdetachmentid, context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// //delete user
// func (c *scoutController) Delete(context *gin.Context) {

// 	rol, _ := middleware.GetRol(c.jwt, context)
// 	if rol > 0 {
// 		c.user.Delete(context)
// 		return
// 	}
// 	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

// }
