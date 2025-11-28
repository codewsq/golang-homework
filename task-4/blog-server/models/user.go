package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);uniqueIndex;not null" json:"username"`
	Email    string `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
}

// HashPassword 加密密码
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
