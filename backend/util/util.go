package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserFromContext(c *gin.Context) (uuid.UUID, error) {
	owner := c.GetString("user")
	if owner != "" {
		ownerId, err := uuid.Parse(owner)
		if err != nil {
			return uuid.New(), err
		}
		return ownerId, nil
	}
	return uuid.New(), errors.New("user not found")
}
