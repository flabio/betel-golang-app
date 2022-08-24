package main

import (
	"bete/Infrastructure/routers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/rol/", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// main
func main() {

	routers.NewRouter()

}
