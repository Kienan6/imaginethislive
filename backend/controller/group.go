package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"itl/model"
	"itl/service"
	"itl/util"
	"net/http"
)

type GroupRoutesController interface {
	createGroup(c *gin.Context)
	getGroup(c *gin.Context)
	getOwnedGroups(c *gin.Context)
	getUsers(c *gin.Context)
	getPosts(c *gin.Context)
}

type GroupRoutesControllerImpl struct {
	GroupService service.GroupService
}

type GroupRoutesParams struct {
	fx.In
	GroupService service.GroupService
}

func (controller *GroupRoutesControllerImpl) createGroup(c *gin.Context) {
	var group model.Group

	if c.ShouldBind(&group) == nil {
		owner, err := util.GetUserFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user required",
			})
			return
		}
		group.OwnerID = owner
		err = controller.GroupService.CreateGroup(&group)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, group)
}

func (controller *GroupRoutesControllerImpl) getGroup(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	group, err := controller.GroupService.GetGroup(idParsed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (controller *GroupRoutesControllerImpl) getOwnedGroups(c *gin.Context) {
	owner, err := util.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user required",
		})
		return
	}
	groups, err := controller.GroupService.FindByOwner(owner)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (controller *GroupRoutesControllerImpl) getUsers(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	groups, err := controller.GroupService.GetUsers(idParsed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (controller *GroupRoutesControllerImpl) getPosts(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	posts, err := controller.GroupService.GetPosts(idParsed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// NewGroupRoutesController
// Setup Group controller
func NewGroupRoutesController(params GroupRoutesParams) Controller {

	controller := &GroupRoutesControllerImpl{
		GroupService: params.GroupService,
	}
	return Controller{
		Group: func(rg *gin.RouterGroup) *gin.RouterGroup {
			group := rg.Group("/group")
			group.POST("/create", controller.createGroup)
			group.GET("/:id/users", controller.getUsers)
			group.GET("/:id/posts", controller.getPosts)
			group.GET("/:id", controller.getGroup)
			group.GET("/groups", controller.getOwnedGroups)
			return group
		},
	}
}
