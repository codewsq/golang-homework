package api

import (
	"github.com/codewsq/blog/server/database"
	"github.com/codewsq/blog/server/logger"
	"github.com/codewsq/blog/server/middleware"
	"github.com/codewsq/blog/server/models"
	"github.com/codewsq/blog/server/responses"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthApi struct{}

// Register 用户注册
func (ac *AuthApi) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required,min=3,max=50"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"input": input,
		}).Warn("Register request validation failed")
		responses.BadRequest(c, err.Error())
		return
	}

	// 检查用户是否已存在
	var existingUser models.User
	result := database.GetDB().Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser)
	if result.Error == nil {
		logger.WithFields(logrus.Fields{
			"username": input.Username,
			"email":    input.Email,
		}).Warn("Register attempt with existing username or email")
		responses.BadRequest(c, "Username or email already exists")
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
	}

	// 加密密码
	if err := user.HashPassword(input.Password); err != nil {
		logger.WithFields(logrus.Fields{
			"username": input.Username,
			"error":    err.Error(),
		}).Error("Failed to hash password")
		responses.InternalServerError(c, "Could not process password")
		return
	}

	// 创建用户
	if err := database.GetDB().Create(&user).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"username": input.Username,
			"error":    err.Error(),
		}).Error("Failed to create user")
		responses.InternalServerError(c, "Could not create user")
		return
	}

	logger.WithFields(logrus.Fields{
		"user_id":  user.ID,
		"username": user.Username,
	}).Info("User registered successfully")

	responses.Created(c, "User registered successfully", gin.H{
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Login 用户登录
func (ac *AuthApi) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Login request validation failed")
		responses.BadRequest(c, err.Error())
		return
	}

	// 查找用户
	var user models.User
	if err := database.GetDB().Where("username = ?", input.Username).First(&user).Error; err != nil {
		logger.WithFields(logrus.Fields{
			"username": input.Username,
			"error":    "user not found",
		}).Warn("Login attempt for non-existent user")
		responses.Unauthorized(c, "Invalid credentials")
		return
	}

	// 验证密码
	if err := user.CheckPassword(input.Password); err != nil {
		logger.WithFields(logrus.Fields{
			"username": input.Username,
			"error":    "invalid password",
		}).Warn("Login attempt with invalid password")
		responses.Unauthorized(c, "Invalid credentials")
		return
	}

	// 生成JWT token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"user_id": user.ID,
			"error":   err.Error(),
		}).Error("Failed to generate JWT token")
		responses.InternalServerError(c, "Could not generate authentication token")
		return
	}

	logger.WithFields(logrus.Fields{
		"user_id":  user.ID,
		"username": user.Username,
	}).Info("User logged in successfully")

	responses.Success(c, "Login successful", gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
