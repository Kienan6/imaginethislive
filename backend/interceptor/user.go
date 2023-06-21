package interceptor

import (
	"github.com/gin-gonic/gin"
	"log"
)

type UserMiddlewareImpl struct {
}

func (mw *UserMiddlewareImpl) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		if len(header) > 0 {
			user := c.Request.Header["Authorization"][0]
			log.Printf("user found %s", user)
			c.Set("user", user)
		} else {
			log.Printf("user not found")
		}
		c.Next()
	}
}

func NewUserMiddleware() Middleware {
	return &UserMiddlewareImpl{}
}
