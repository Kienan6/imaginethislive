package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"itl/model"
	"itl/service"
	"itl/util"
	"net/http"
)

type UserRoutesController interface {
	create(c *gin.Context)
	getGroups(c *gin.Context)
}

type UserRoutesControllerImpl struct {
	UserService service.UserService
}

type UserRoutesParams struct {
	fx.In
	UserService service.UserService
}

func (controller *UserRoutesControllerImpl) create(c *gin.Context) {
	var user model.User

	if c.ShouldBind(&user) == nil {
		err := controller.UserService.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserRoutesControllerImpl) getGroups(c *gin.Context) {
	id, err := util.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	groups, err := controller.UserService.GetGroups(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// NewUserRoutesController
// Setup User controller
func NewUserRoutesController(params UserRoutesParams) Controller {

	controller := &UserRoutesControllerImpl{
		UserService: params.UserService,
	}
	return Controller{
		Group: func(rg *gin.RouterGroup) *gin.RouterGroup {
			user := rg.Group("/user")
			user.POST("/create", controller.create)
			user.GET("/groups", controller.getGroups)
			return user
		},
	}
}
