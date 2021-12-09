package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/UseCases/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

//KanbanController is a ....
type KanbanController interface {
	GetKanban(context *gin.Context)
	CountKanban(context *gin.Context)
}

type kanbanController struct {
	kanban services.KanbanService
	jwt    services.JWTService
}

//NewKanbanController is creating anew instance of UserControlller
func NewKanbanController() KanbanController {

	return &kanbanController{
		kanban: services.NewKanbanService(),
		jwt:    services.NewJWTService(),
	}
}

func (c *kanbanController) GetKanban(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.kanban.GetKanbans(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

func (c *kanbanController) CountKanban(context *gin.Context) {

	claim := middleware.GetRol(c.jwt, context)
	if claim.Rol > 0 {
		c.kanban.GetCountKanbans(context)

		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
