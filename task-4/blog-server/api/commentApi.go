package api

import (
	"github.com/codewsq/blog/server/database"
	"github.com/codewsq/blog/server/models"
	"github.com/codewsq/blog/server/responses"
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
		responses.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := c.GetUint("userID")

	comment := models.Comment{
		Content: input.Content,
		UserID:  userID,
		PostID:  uint(input.PostID),
	}
	if err := database.GetDB().Create(&comment).Error; err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	// 查询评论所属的文章信息
	var createComment models.Comment
	database.GetDB().Preload("Post.User").Preload("User").First(&createComment, comment.ID)
	responses.Created(c, "comment created successfully", createComment)
}

func (api *CommentApi) GetComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		responses.BadRequest(c, "Invalid post ID")
	}
	var comment models.Comment
	if err := database.GetDB().Preload("Post.User").Where(&models.Comment{PostID: uint(postId)}).
		Find(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
		responses.Error(c, http.StatusBadRequest, "Comment not found")
	}
	responses.Success(c, "comment successfully retrieved", comment)
}
