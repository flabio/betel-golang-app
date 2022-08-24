package tests

import (
	"bete/UseCases/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// service of all
func TestRolAll(t *testing.T) {
	c := require.New(t)
	var context *gin.Context
	result := services.NewRolService()
	err := result.GetAllService(context)
	if err != nil {
		t.Fail()
	}
	//d.Equal(expect, result)
}
