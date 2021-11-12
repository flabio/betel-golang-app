package controllers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"fmt"

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

	rol, _ := middleware.GetRol(c.jwt, context)
	fmt.Println(rol)
	c.kanban.GetKanbans(context)

}

func (c *kanbanController) CountKanban(context *gin.Context) {

	rol, _ := middleware.GetRol(c.jwt, context)
	fmt.Println(rol)
	c.kanban.GetCountKanbans(context)

	// }
	// context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
