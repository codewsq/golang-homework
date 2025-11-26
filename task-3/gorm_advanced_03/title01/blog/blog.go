package blog

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	PostCount int
	Posts     []Post `gorm:"foreignkey:UserID"`
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint
	CommentStatus string
	Comments      []Comment `gorm:"foreignkey:PostID"`
}

type Comment struct {
	gorm.Model
	Context string
	PostID  uint
}

func CreateTable(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	return err
}

func SaveInfo(db *gorm.DB) error {
	tx := db.Begin()
	user := User{Name: "张三", PostCount: 2}
	err := tx.Create(&user).Error
	if err != nil {
		fmt.Println("插入用户信息失败")
		tx.Rollback()
		return err
	}

	posts := []Post{
		{Title: "gorm指南", Content: "指南概述内容", UserID: user.ID},
		{Title: "java入门到跑路", Content: "java spring框架内容", UserID: user.ID},
	}
	createPost := tx.Create(&posts)
	if createPost.Error != nil {
		fmt.Println("插入文章信息失败")
		tx.Rollback()
		return createPost.Error
	}
	fmt.Println("插入文章信息数：", createPost.RowsAffected)

	comments := make([]Comment, 0)
	comments = append(comments, Comment{Context: posts[0].Title + "-写的不错", PostID: posts[0].ID})
	comments = append(comments, Comment{Context: posts[0].Title + "-通俗易懂", PostID: posts[0].ID})
	comments = append(comments, Comment{Context: posts[1].Title + "-根本学不完", PostID: posts[1].ID})

	createCom := tx.Create(&comments)
	if createCom.Error != nil {
		fmt.Println("插入评价信息失败")
		tx.Rollback()
		return createCom.Error
	}
	fmt.Println("插入评论信息数：", createCom.RowsAffected)
	return tx.Commit().Error
}
