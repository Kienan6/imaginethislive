package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return true
	}
	return false
}
