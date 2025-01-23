package rest

import (
	"github.com/google/uuid"
	"time"
)
type Comment struct {
	ID uuid.UUID
	UserID uuid.UUID
	PostID uuid.UUID
	Text string
	CreatedAt time.Time
}
	