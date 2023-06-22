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

type PostRoutesController interface {
	createPost(c *gin.Context)
	getPost(c *gin.Context)
}

type PostRoutesControllerImpl struct {
	PostService service.PostService
}

type PostRoutesParams struct {
	fx.In
	PostService service.PostService
}

func (controller *PostRoutesControllerImpl) createPost(c *gin.Context) {
	var Post model.Post

	if c.ShouldBind(&Post) == nil {
		owner, err := util.GetUserFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user required",
			})
			return
		}
		Post.UserID = owner
		err = controller.PostService.CreatePost(&Post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, Post)
}

func (controller *PostRoutesControllerImpl) getPost(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	Post, err := controller.PostService.GetPost(idParsed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Post)
}

// NewPostRoutesController
// Setup Post controller
func NewPostRoutesController(params PostRoutesParams) Controller {

	controller := &PostRoutesControllerImpl{
		PostService: params.PostService,
	}
	return Controller{
		Group: func(rg *gin.RouterGroup) *gin.RouterGroup {
			post := rg.Group("/post")
			post.POST("/create", controller.createPost)
			post.GET("/:id", controller.getPost)
			return post
		},
	}
}
