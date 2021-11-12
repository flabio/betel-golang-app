package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"

	"net/http"

	"github.com/gin-gonic/gin"
)

//UserService is a contract.....
type KanbanService interface {
	GetKanbans(context *gin.Context)
	GetCountKanbans(context *gin.Context)
}

type kanbanService struct {
	kanbanR repositorys.UserRepository
}

//NewUserService creates a new instance of UserService
func NewKanbanService() KanbanService {
	var kanban repositorys.UserRepository = repositorys.NewUserRepository()
	return &kanbanService{
		kanbanR: kanban,
	}
}

func (kanbanService *kanbanService) GetKanbans(context *gin.Context) {

	navigatores, err := kanbanService.kanbanR.ListNavigators()
	pioneers, err := kanbanService.kanbanR.ListPioneers()
	followers, err := kanbanService.kanbanR.ListFollowersWays()
	scouts, err := kanbanService.kanbanR.ListScouts()
	commanders, err := kanbanService.kanbanR.ListCommanders()
	if err != nil {
		validadErrors(err, context)
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

	count_navigators, count_pioneers, count_followers, count_scouts, count_commanders := kanbanService.kanbanR.CounKanban()

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
