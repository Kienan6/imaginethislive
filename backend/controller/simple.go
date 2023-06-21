package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"itl/service"
	"net/http"
)

type SimpleRoutesController interface {
	testHandler(c *gin.Context)
}

type SimpleRoutesControllerImpl struct {
	SimpleService service.SimpleService
}

type SimpleRoutesParams struct {
	fx.In
	SimpleService service.SimpleService
}

func (controller *SimpleRoutesControllerImpl) testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": controller.SimpleService.GetText(),
	})
}

// NewSimpleRoutesController
// Setup simple controller
func NewSimpleRoutesController(params SimpleRoutesParams) Controller {

	controller := &SimpleRoutesControllerImpl{
		SimpleService: params.SimpleService,
	}
	return Controller{
		Group: func(rg *gin.RouterGroup) *gin.RouterGroup {
			simple := rg.Group("/another")
			simple.GET("/test", controller.testHandler)
			return simple
		},
	}
}
