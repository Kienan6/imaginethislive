package interceptor

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type UserMiddlewareImpl struct {
}

func parseBasicAuth(header string) (username, password string, ok bool) {
	s := strings.SplitN(header, " ", 2)
	if len(s) != 2 {
		return "", "", false
	}
	decoded, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return "", "", false
	}
	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return "", "", false
	}
	return credentials[0], credentials[1], true

}

// TODO: Documentation
// Basic Auth
func (mw *UserMiddlewareImpl) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		if len(header) > 0 {
			auth := c.Request.Header["Authorization"][0]

			user, _, ok := parseBasicAuth(auth)
			if ok {
				log.Printf("user found - %s", user)
				c.Set("user", user)
			} else {
				log.Print("user not found")
			}
		} else {
			log.Printf("authorization header not found")
		}

		c.Next()
	}
}

func NewUserMiddleware() Middleware {
	return &UserMiddlewareImpl{}
}
