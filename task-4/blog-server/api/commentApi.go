package api

import (
	"github.com/codewsq/blog/server/database"
	"github.com/codewsq/blog/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentApi struct{}

func (api *CommentApi) CreateComment(c *gin.Context) {
	type Input struct {
		Content string `json:"content" binding:"required"`
		PostID  int    `json:"post_id" binding:"required"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("userID")

	comment := models.Comment{
		Content: input.Content,
		UserID:  userID,
		PostID:  uint(input.PostID),
	}
	if err := database.GetDB().Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 查询评论所属的文章信息
	var createComment models.Comment
	database.GetDB().Preload("Post.User").Preload("User").First(&createComment, comment.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "comment created successfully",
		"comment": createComment,
	})
}

func (api *CommentApi) GetComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
	}
	var comment models.Comment
	if err := database.GetDB().Preload("Post.User").Where(&models.Comment{PostID: uint(postId)}).
		Find(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
	}
	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

/*
type Comment struct {
	gorm.Model
	Content   string    `gorm:"not null"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID" json:"post,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
*/
