package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	Username  string
	Groups    []*Group  `gorm:"many2many:user_group;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
}
