package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Group struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	OwnerID   uuid.UUID
	Name      string
	Users     []*User `gorm:"many2many:user_group;"`
	Posts     []Post
	CreatedAt pgtype.Timestamp
	Owner     User
}
