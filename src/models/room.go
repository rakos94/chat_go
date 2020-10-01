package models

import (
	"time"
)

// Room models
type Room struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Rooms models
type Rooms []Room
