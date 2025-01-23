package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID          uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	UserID      uuid.UUID
	GroupID     uuid.UUID
	Uri         string
	Description string
	Plays       int
	Upvotes     int
	Downvotes   int
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	User        User
	Comments    []Comment
}
