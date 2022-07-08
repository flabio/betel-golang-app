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
	user InterfacesService.IUserService
	role InterfacesService.IUserRolService
	jwt  InterfacesService.IJWTService
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

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.user.SetCreateService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//update profile method pust
func (c *userController) UpdateProfile(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.user.SetUpdateService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//update password
func (c *userController) PasswordChange(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.user.SetUpdatePasswordService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//update user method push
func (c *userController) Update(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.user.SetUpdateService(context)

		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))

}

func (c *userController) ListUser(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.user.GetListUserService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

func (c *userController) FindUserNameLastName(context *gin.Context) {
	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.user.GetFindUserNameLastNameService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

// func (c *userController) All(context *gin.Context) {

// 	rol := middleware.ValidadToken(c.jwt, context)
// 	if rol == 1 {
// 		c.user.All(context)
// 		return
// 	}
// 	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest,constantvariables.PERMISSION_DANIED))
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

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.user.SetRemoveService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))

}

//profile user
func (c *userController) Profile(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol == 1 {
		c.user.GetProfileService(uint(claim.Id), context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))
}

//profile user
func (c *userController) FindUser(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.user.GetFindUserService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(http.StatusBadRequest, constantvariables.PERMISSION_DANIED))

}
