package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
	CreatedAt   pgtype.Timestamp
	User        User
}
