package api

import (
	"github.com/codewsq/blog/server/database"
	"github.com/codewsq/blog/server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostApi struct{}

// CreatePost 创建文章
func (pc *PostApi) CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}

	if err := database.GetDB().Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create post"})
		return
	}

	// 查询创建的文章（包含用户信息）
	var createdPost models.Post
	database.GetDB().Preload("User").First(&createdPost, post.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    createdPost,
	})
}

// GetPosts 获取文章列表
func (pc *PostApi) GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := database.GetDB().Preload("User").Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// GetPost 获取单篇文章
func (pc *PostApi) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := database.GetDB().Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// UpdatePost 更新文章
func (pc *PostApi) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")

	// 检查文章是否存在且属于当前用户
	var post models.Post
	if err := database.GetDB().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own posts"})
		return
	}

	// 更新文章
	updates := map[string]interface{}{}
	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Content != "" {
		updates["content"] = input.Content
	}

	if err := database.GetDB().Model(&post).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update post"})
		return
	}

	// 重新查询更新后的文章
	database.GetDB().Preload("User").First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}

// DeletePost 删除文章
func (pc *PostApi) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	userID := c.GetUint("userID")

	// 检查文章是否存在且属于当前用户
	var post models.Post
	if err := database.GetDB().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own posts"})
		return
	}

	if err := database.GetDB().Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post deleted successfully",
		"post":    post,
	})
}
