package api

import (
	"github.com/codewsq/blog/server/database"
	"github.com/codewsq/blog/server/logger"
	"github.com/codewsq/blog/server/models"
	"github.com/codewsq/blog/server/responses"
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostApi struct{}

// CreatePost 创建文章
func (pc *PostApi) CreatePost(c *gin.Context) {
	userID := c.GetUint("userID")

	var input struct {
		Title   string `json:"title" binding:"required,min=1,max=255"`
		Content string `json:"content" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err.Error(),
		}).Warn("Create post request validation failed")
		responses.BadRequest(c, err.Error())
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}

	if err := database.GetDB().Create(&post).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err.Error(),
		}).Error("Failed to create post")
		responses.InternalServerError(c, "Could not create post")
		return
	}

	// 查询创建的文章（包含用户信息）
	var createdPost models.Post
	database.GetDB().Preload("User").First(&createdPost, post.ID)

	logger.WithFields(logrus.Fields{
		"post_id": post.ID,
		"user_id": userID,
	}).Info("Post created successfully")

	responses.Created(c, "Post created successfully", gin.H{
		"post": createdPost,
	})
}

// GetPosts 获取文章列表
func (pc *PostApi) GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := database.GetDB().Preload("User").Order("created_at desc").Find(&posts).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to fetch posts")
		responses.InternalServerError(c, "Could not fetch posts")
		return
	}

	logger.WithFields(logrus.Fields{
		"count": len(posts),
	}).Debug("Posts fetched successfully")

	responses.Success(c, "Posts fetched successfully", gin.H{
		"posts": posts,
	})
}

// GetPost 获取单篇文章
func (pc *PostApi) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.WithFields(logrus.Fields{
			"post_id": c.Param("id"),
			"error":   "invalid post id format",
		}).Warn("Invalid post ID")
		responses.BadRequest(c, "Invalid post ID")
		return
	}

	var post models.Post
	if err := database.GetDB().Preload("User").First(&post, id).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"post_id": id,
			"error":   err.Error(),
		}).Warn("Post not found")
		responses.NotFound(c, "Post not found")
		return
	}

	logger.WithFields(logrus.Fields{
		"post_id": id,
	}).Debug("Post fetched successfully")

	responses.Success(c, "Post fetched successfully", gin.H{
		"post": post,
	})
}

// UpdatePost 更新文章
func (pc *PostApi) UpdatePost(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.WithFields(logrus.Fields{
			"post_id": c.Param("id"),
			"error":   "invalid post id format",
		}).Warn("Invalid post ID")
		responses.BadRequest(c, "Invalid post ID")
		return
	}

	var input struct {
		Title   string `json:"title" binding:"omitempty,min=1,max=255"`
		Content string `json:"content" binding:"omitempty,min=1"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.WithFields(logrus.Fields{
			"user_id": userID,
			"post_id": id,
			"error":   err.Error(),
		}).Warn("Update post request validation failed")
		responses.BadRequest(c, err.Error())
		return
	}

	// 检查文章是否存在且属于当前用户
	var post models.Post
	if err := database.GetDB().First(&post, id).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"post_id": id,
			"error":   err.Error(),
		}).Warn("Post not found for update")
		responses.NotFound(c, "Post not found")
		return
	}

	if post.UserID != userID {
		logger.WithFields(logrus.Fields{
			"user_id":    userID,
			"post_id":    id,
			"post_owner": post.UserID,
		}).Warn("Attempt to update another user's post")
		responses.Forbidden(c, "You can only update your own posts")
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

	if len(updates) == 0 {
		logger.WithFields(logrus.Fields{
			"user_id": userID,
			"post_id": id,
		}).Warn("Update post with no valid fields")
		responses.BadRequest(c, "No valid fields to update")
		return
	}

	if err := database.GetDB().Model(&post).Updates(updates).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"user_id": userID,
			"post_id": id,
			"error":   err.Error(),
		}).Error("Failed to update post")
		responses.InternalServerError(c, "Could not update post")
		return
	}

	// 重新查询更新后的文章
	database.GetDB().Preload("User").First(&post, id)

	logger.WithFields(logrus.Fields{
		"post_id": id,
		"user_id": userID,
	}).Info("Post updated successfully")

	responses.Success(c, "Post updated successfully", gin.H{
		"post": post,
	})
}

// DeletePost 删除文章
func (pc *PostApi) DeletePost(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.WithFields(logrus.Fields{
			"post_id": c.Param("id"),
			"error":   "invalid post id format",
		}).Warn("Invalid post ID")
		responses.BadRequest(c, "Invalid post ID")
		return
	}

	// 检查文章是否存在且属于当前用户
	var post models.Post
	if err := database.GetDB().First(&post, id).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"post_id": id,
			"error":   err.Error(),
		}).Warn("Post not found for deletion")
		responses.NotFound(c, "Post not found")
		return
	}

	if post.UserID != userID {
		logger.WithFields(logrus.Fields{
			"user_id":    userID,
			"post_id":    id,
			"post_owner": post.UserID,
		}).Warn("Attempt to delete another user's post")
		responses.Forbidden(c, "You can only delete your own posts")
		return
	}

	if err := database.GetDB().Delete(&post).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"user_id": userID,
			"post_id": id,
			"error":   err.Error(),
		}).Error("Failed to delete post")
		responses.InternalServerError(c, "Could not delete post")
		return
	}

	logger.WithFields(logrus.Fields{
		"post_id": id,
		"user_id": userID,
	}).Info("Post deleted successfully")

	responses.Success(c, "Post deleted successfully", nil)
}
