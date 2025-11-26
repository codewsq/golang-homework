package blog

import (
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
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
*/
// 方式一
func SelectInfoByUserId(db *gorm.DB, userID uint) (*User, error) {
	posts := []Post{}
	db.Where(&Post{UserID: userID}).Find(&posts)

	for i := range posts {
		comments := []Comment{}
		db.Where(&Comment{PostID: posts[i].ID}).Find(&comments)
		posts[i].Comments = comments
	}

	user := User{}
	db.First(&user, userID)
	user.Posts = posts
	return &user, nil
}

// 方式二（推荐）
func SelectInfoByUserId2(db *gorm.DB, userID uint) (*User, error) {
	var user User
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

/*
编写Go代码，使用Gorm查询评论数量最多的文章信息
*/
func SelectPostByMaxCom(db *gorm.DB) (*Post, error) {
	var post Post
	err := db.Select("posts.*,count(c.id) as comment_count").
		Joins("left join comments c on posts.id = c.post_id").
		Group("posts.id").
		Order("comment_count desc").
		First(&post).Error
	return &post, err
}
