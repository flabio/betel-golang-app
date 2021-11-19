package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubdetachmentController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
	FindByIdDetachment(context *gin.Context)
	AddUserSubDetachment(context *gin.Context)
	RemoveUserSubDetachment(context *gin.Context)
}

type subdetachmentController struct {
	subDetachment services.SubDetachmentService
	jwt           services.JWTService
}

func NewSubdetachmentController() SubdetachmentController {

	return &subdetachmentController{
		subDetachment: services.NewSubDetachmentService(),
		jwt:           services.NewJWTService(),
	}
}

//GET /subdetachment
// get list of subdetachment
func (c *subdetachmentController) All(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.All(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// get
func (c *subdetachmentController) FindById(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.FindById(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//FindByIdDetachment
func (c *subdetachmentController) FindByIdDetachment(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.FindByIdDetachment(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//create subdetachment method post
func (c *subdetachmentController) Create(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.Create(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//update subdetachment method put
func (c *subdetachmentController) Update(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.Update(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// delete subdetachment
func (c *subdetachmentController) Remove(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.Remove(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

//create AddUserSubDetachment method post
func (c *subdetachmentController) AddUserSubDetachment(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.AddUserSubDetachment(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//create AddUserSubDetachment method post
func (c *subdetachmentController) RemoveUserSubDetachment(context *gin.Context) {
	rol, _ := middleware.GetRol(c.jwt, context)
	if rol == 1 {
		c.subDetachment.RemoveUserSubDetachment(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
