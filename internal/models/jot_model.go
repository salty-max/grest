package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jot struct {
	gorm.Model
	UUID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Slug   string    `gorm:"not null"`
	Body   string    `gorm:"not null"`
	Author string    `gorm:"not null"`
}
