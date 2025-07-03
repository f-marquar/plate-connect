package db

import (
	"time"
)

type User struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Kennzeichen string    `gorm:"uniqueIndex;size:20;not null"`
	Email       string    `gorm:"size:100;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
