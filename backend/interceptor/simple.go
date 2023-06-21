package interceptor

import (
	"github.com/gin-gonic/gin"
	"log"
)

type SimpleMiddlewareImpl struct {
}

func (mw *SimpleMiddlewareImpl) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("From the simple middleware: %d", c.Writer.Status())
		c.Next()
	}
}

func NewSimpleMiddleware() Middleware {
	return &SimpleMiddlewareImpl{}
}
