package controllers

import (
	"bete/Infrastructure/middleware"

	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserController is a ....
type UserController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	UpdateProfile(context *gin.Context)
	PasswordChange(context *gin.Context)
	//UsersRoles(context *gin.Context)
	Delete(context *gin.Context)

	Profile(context *gin.Context)
	FindUser(context *gin.Context)
	//All(context *gin.Context)
	FindUserNameLastName(context *gin.Context)
	ListUser(context *gin.Context)
}

type userController struct {
	user services.UserService
	role services.UserRolService
	jwt  services.JWTService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController() UserController {
	return &userController{
		user: services.NewUserService(),
		role: services.NewUserRolService(),
		jwt:  services.NewJWTService(),
	}
}

//create user method post
func (c *userController) Create(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.user.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//update profile method pust
func (c *userController) UpdateProfile(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.user.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//update password
func (c *userController) PasswordChange(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.user.UpdatePassword(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//update user method push
func (c *userController) Update(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.user.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

func (c *userController) ListUser(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)

	if rol == 1 {
		c.user.ListUser(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

func (c *userController) FindUserNameLastName(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)

	if rol > 0 {
		c.user.FindUserNameLastName(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// func (c *userController) All(context *gin.Context) {

// 	rol := middleware.GetRol(c.jwt, context)
// 	if rol == 1 {
// 		c.user.All(context)
// 		return
// 	}
// 	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
// 	c.user.All(context)
// }

//role
// func (c *userController) UsersRoles(context *gin.Context) {
// 	var users []entity.Role = c.role.AllUserRole()
// 	res := utilities.BuildResponse(true, "OK", users)
// 	context.JSON(http.StatusOK, res)
// }

//delete user
func (c *userController) Delete(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.user.Delete(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

//profile user
func (c *userController) Profile(context *gin.Context) {

	rol, id := middleware.GetRol(c.jwt, context)
	if rol == 1 {

		c.user.Profile(uint(id), context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//profile user
func (c *userController) FindUser(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol > 0 {

		c.user.FindUser(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}
