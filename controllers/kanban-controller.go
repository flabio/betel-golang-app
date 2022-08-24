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

// KanbanController is a ....
type KanbanController interface {
	GetKanban(context *gin.Context)
	CountKanban(context *gin.Context)
}

type kanbanController struct {
	kanban InterfacesService.IKanbanService
	jwt    InterfacesService.IJWTService
}

// NewKanbanController is creating anew instance of UserControlller
func NewKanbanController() KanbanController {

	return &kanbanController{
		kanban: services.NewKanbanService(),
		jwt:    services.NewJWTService(),
	}
}

func (c *kanbanController) GetKanban(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.kanban.GetKanbans(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}

func (c *kanbanController) CountKanban(context *gin.Context) {

	claim := middleware.ValidadToken(c.jwt, context)
	if claim.Rol > 0 {
		c.kanban.GetCountKanbans(context)

		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildErrResponse(constantvariables.PERMISSION_DANIED))
}
