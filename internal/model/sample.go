package model

import (
	"time"

	"gorm.io/gorm"
)

type Sample struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey;" json:"id"`
	Title       string    `json:"title"`
	Description int       `json:"description"`
	Phone       int       `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}
