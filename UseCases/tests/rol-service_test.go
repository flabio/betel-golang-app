package tests

import (
	"bete/UseCases/services"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestRolAllSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	services.NewRolService().GetFindByIdService(c)
	assert.Equal(t, 200, w.Code) // or what value you need it to be

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, 1) // want is a gin.H that contains the

	// gin.SetMode(gin.TestMode)
	// t.Run("success", func(t *testing.T) {
	// 	router := gin.Default()
	// 	router.GET("api/v1/rol")
	// 	router.Run()

	// })
	// var context *gin.Context
	// r := rolTestController{}
	// r.rol.GetAllService(context)
	// if r.rol.GetAllService(context) != expect {
	// 	t.Errorf("got %q, expected %d", r.rol.GetAllService(context), expect)
	// }
}
