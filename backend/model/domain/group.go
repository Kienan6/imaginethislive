package domain

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	OwnerID   uuid.UUID
	Name      string
	Users     []*User `gorm:"many2many:user_group;"`
	Posts     []Post
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	Owner     User
}
