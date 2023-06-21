package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	Username  string
	Groups    []*Group `gorm:"many2many:user_group;"`
	CreatedAt pgtype.Timestamp
}
