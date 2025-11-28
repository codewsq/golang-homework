package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Published bool      `gorm:"default:false" json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
