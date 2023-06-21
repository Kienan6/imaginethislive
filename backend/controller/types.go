package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Controller struct {
	fx.Out
	Group func(rg *gin.RouterGroup) *gin.RouterGroup `group:"controllers"`
}
