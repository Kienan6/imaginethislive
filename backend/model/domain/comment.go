package domain

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()" rest:"include"`
	UserID    uuid.UUID `rest:"include"`
	PostID    uuid.UUID `rest:"include"`
	Text      string    `rest:"include"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" rest:"include"`
	User      User
	Post      Post
}
