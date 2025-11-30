package middleware

import (
	"github.com/codewsq/blog/server/logger"
	"github.com/codewsq/blog/server/responses"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录panic信息
				logger.Errorf("Panic recovered: %v\n%s", err, string(debug.Stack()))

				// 返回500错误
				responses.InternalServerError(c, "Internal Server Error")
				c.Abort()
			}
		}()

		c.Next()

		// 处理其他错误
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logger.Errorf("Request error: %v", e)
			}

			// 如果没有设置状态码，默认返回500
			if c.Writer.Status() == http.StatusOK {
				responses.InternalServerError(c, "Internal Server Error")
			}
		}
	}
}

// RecoveryMiddleware 恢复中间件（记录更多上下文信息）
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录详细的错误信息
				logger.WithFields(logrus.Fields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
					"ip":     c.ClientIP(),
					"panic":  err,
					"stack":  string(debug.Stack()),
				}).Error("Recovered from panic")

				responses.InternalServerError(c, "Internal Server Error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
