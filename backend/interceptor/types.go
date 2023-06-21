package interceptor

import (
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Run() gin.HandlerFunc
}
