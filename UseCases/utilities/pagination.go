package utilities

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context, limit int) (int, int) {
	p := c.Query("page")

	if p == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		return 1, 0
	}

	begin := (limit * page) - limit
	return page, begin
}
