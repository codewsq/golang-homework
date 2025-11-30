package middleware

import (
	"github.com/codewsq/blog/server/config"
	"github.com/codewsq/blog/server/models"
	"github.com/codewsq/blog/server/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("your-secret-key") // 生产环境请使用环境变量

// Claims JWT声明结构
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(user models.User) (string, error) {
	cfg := config.GetConfig()

	expirationTime := time.Now().Add(time.Duration(cfg.JWT.ExpirationHours) * time.Hour)

	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    cfg.App.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// JWTMiddleware JWT认证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.Unauthorized(c, "Authorization header is required")
			c.Abort() // 立即终止当前请求的处理链，阻止后续的中间件和处理函数被执行
			return
		}

		// 检查Bearer token格式
		/*parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}*/

		tokenString := authHeader
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			responses.Unauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
