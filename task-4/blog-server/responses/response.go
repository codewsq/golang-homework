package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// Created 创建成功响应
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Code:    http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: "error",
		Error:   errorMessage,
	})
}

// BadRequest 400错误
func BadRequest(c *gin.Context, errorMessage string) {
	Error(c, http.StatusBadRequest, errorMessage)
}

// Unauthorized 401错误
func Unauthorized(c *gin.Context, errorMessage string) {
	Error(c, http.StatusUnauthorized, errorMessage)
}

// Forbidden 403错误
func Forbidden(c *gin.Context, errorMessage string) {
	Error(c, http.StatusForbidden, errorMessage)
}

// NotFound 404错误
func NotFound(c *gin.Context, errorMessage string) {
	Error(c, http.StatusNotFound, errorMessage)
}

// InternalServerError 500错误
func InternalServerError(c *gin.Context, errorMessage string) {
	Error(c, http.StatusInternalServerError, errorMessage)
}
