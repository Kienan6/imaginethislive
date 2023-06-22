package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID
	PostID    uuid.UUID
	Text      string
	CreatedAt pgtype.Timestamp
	User      User
	Post      Post
}
