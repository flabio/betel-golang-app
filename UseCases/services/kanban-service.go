package services

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/utilities"

	"net/http"

	"github.com/gin-gonic/gin"
)

type kanbanService struct {
	kanbanR Interfaces.IUser
}

//NewUserService creates a new instance of UserService
func NewKanbanService() InterfacesService.IKanbanService {
	kanban := repositorys.NewUserRepository()
	return &kanbanService{
		kanbanR: kanban,
	}
}

func (kanbanService *kanbanService) GetKanbans(context *gin.Context) {

	navigatores, err := kanbanService.kanbanR.GetListNavigators()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	pioneers, err := kanbanService.kanbanR.GetListPioneers()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	followers, err := kanbanService.kanbanR.GetListFollowersWays()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	scouts, err := kanbanService.kanbanR.GetListScouts()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	commanders, err := kanbanService.kanbanR.GetListCommanders()
	if err != nil {
		res := utilities.BuildErrResponse(http.StatusBadRequest, err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	context.JSON(http.StatusOK, struct {
		GetNavigators    []entity.User `json:"navigatores"`
		GetPioneers      []entity.User `json:"pioneers"`
		GetFollowersWays []entity.User `json:"followers"`
		GetScouts        []entity.User `json:"scouts"`
		GetCommanders    []entity.User `json:"commanders"`
	}{
		GetNavigators:    navigatores,
		GetPioneers:      pioneers,
		GetFollowersWays: followers,
		GetScouts:        scouts,
		GetCommanders:    commanders,
	})

}

func (kanbanService *kanbanService) GetCountKanbans(context *gin.Context) {

	count_navigators, count_pioneers, count_followers, count_scouts, count_commanders := kanbanService.kanbanR.GetCounKanban()

	context.JSON(http.StatusOK, struct {
		CountNavigators    int64 `json:"count_navigators"`
		CountPioneers      int64 `json:"count_pioneers"`
		CountFollowersWays int64 `json:"count_followers"`
		CountScouts        int64 `json:"count_scouts"`
		CountCommanders    int64 `json:"count_commanders"`
	}{
		CountNavigators:    count_navigators,
		CountPioneers:      count_pioneers,
		CountFollowersWays: count_followers,
		CountScouts:        count_scouts,
		CountCommanders:    count_commanders,
	})

}
