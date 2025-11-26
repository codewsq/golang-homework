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

/*
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
*/
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	result := tx.Model(p).Where(&Post{UserID: p.UserID}).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	user := User{}
	tx.First(&user, p.UserID)
	update := tx.Model(&user).Update("PostCount", count)
	if update.Error != nil {
		return update.Error
	}
	fmt.Println("更新用户信息表文章数量：", count)
	return
}

/*
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("调用评论删除钩子。。。", c.PostID)
	var count int64
	result := tx.Model(c).Where(&Comment{PostID: c.PostID}).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return
	}
	post := Post{}
	tx.First(&post, c.PostID)
	update := tx.Model(&post).Update("CommentStatus", "无评论")
	if update.Error != nil {
		return update.Error
	}
	fmt.Println("---> 文章：", post.Title, "，更新评论状态为：无评论")
	return
}

// 新增文章信息
func InsertPost(db *gorm.DB, posts *[]Post) error {
	createPost := db.Create(posts)
	if createPost.Error != nil {
		return createPost.Error
	}
	fmt.Println("插入文章信息数：", createPost.RowsAffected)
	return nil
}

// 删除评论
func DeleteComment(db *gorm.DB, commentID int) error {
	comment := Comment{}
	db.First(&comment, commentID)
	result := db.Delete(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
