package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"itl/model/domain"
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

// createGroup creates a group
//
//	@Summary      Create a Group
//	@Description  create a group
//	@Tags		  group
//	@Accept       json
//	@Produce      json
//	@Param        group    body     model.Group  true "Group Object"
//	@Success      200  {object}   model.Group
//	@Failure      400  {object}	  map[string]any
//	@Security BasicAuth
//	@Router       /group/create [post]
func (controller *GroupRoutesControllerImpl) createGroup(c *gin.Context) {
	var group domain.Group

	if c.ShouldBind(&group) == nil {
		owner, err := util.GetUserFromContext(c)
		if handleError(c, err) {
			return
		}
		group.OwnerID = owner
		groupResp, err := controller.GroupService.CreateGroup(&group)
		if handleError(c, err) {
			return
		}
		c.JSON(http.StatusOK, groupResp)
	}

}

// getGroup gets a group
//
//	@Summary      Create a Group
//	@Description  create a group
//	@Tags		  group
//	@Accept       json
//	@Produce      json
//	@Param        id    path     string  true "Group ID"
//	@Success      200  {object}   model.Group
//	@Failure      400  {object}	  map[string]any
//	@Security BasicAuth
//	@Router       /group/{id} [get]
func (controller *GroupRoutesControllerImpl) getGroup(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := uuid.Parse(id)
	if handleError(c, err) {
		return
	}
	group, err := controller.GroupService.GetGroup(idParsed)
	if handleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, group)
}

func (controller *GroupRoutesControllerImpl) getOwnedGroups(c *gin.Context) {
	owner, err := util.GetUserFromContext(c)
	if handleError(c, err) {
		return
	}
	groups, err := controller.GroupService.FindByOwner(owner)
	if handleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (controller *GroupRoutesControllerImpl) getUsers(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := uuid.Parse(id)
	if handleError(c, err) {
		return
	}
	users, err := controller.GroupService.GetUsers(idParsed)
	if handleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, users)
}

func (controller *GroupRoutesControllerImpl) getPosts(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := uuid.Parse(id)
	if handleError(c, err) {
		return
	}
	posts, err := controller.GroupService.GetPosts(idParsed)
	if handleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, posts)
}

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
