package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Content   string    `gorm:"not null"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID" json:"post,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
