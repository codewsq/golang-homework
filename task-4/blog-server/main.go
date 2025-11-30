package main

import (
	"github.com/codewsq/blog/server/api"
	"github.com/codewsq/blog/server/config"
	"github.com/codewsq/blog/server/database"
	"github.com/codewsq/blog/server/logger"
	"github.com/codewsq/blog/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置文件
	if err := config.LoadConfig(""); err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		panic(err)
	}

	logger.Info("Application starting...")
	logger.WithFields(logrus.Fields{
		"app":     cfg.App.Name,
		"version": cfg.App.Version,
		"mode":    cfg.Server.Mode,
	}).Info("Application configuration loaded")

	// 设置Gin运行模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化数据库连接
	if err := database.ConnectDB(); err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	// 创建Gin路由
	router := gin.Default()

	// 添加中间件
	router.Use(middleware.RecoveryMiddleware()) // 恢复中间件
	router.Use(middleware.ErrorHandler())       // 错误处理中间件
	router.Use(gin.Logger())                    // Gin默认日志

	// 初始化控制器
	authApi := &api.AuthApi{}
	postApi := &api.PostApi{}
	commentApi := &api.CommentApi{}

	// 健康检查路由
	router.GET("/health", func(c *gin.Context) {
		logger.Debug("Health check requested")
		c.JSON(200, gin.H{
			"status":  "ok",
			"app":     cfg.App.Name,
			"version": cfg.App.Version,
		})
	})

	// API路由组
	api := router.Group("/api")
	{
		// 公开路由 - 无需认证
		api.POST("/register", authApi.Register)
		api.POST("/login", authApi.Login)

		// 保护路由 - 需要JWT认证
		protected := api.Group("")
		protected.Use(middleware.JWTMiddleware())
		{
			post := protected.Group("/post")
			{
				post.PUT("/create", postApi.CreatePost)
				post.POST("/update/:id", postApi.UpdatePost)
				post.DELETE("/delete/:id", postApi.DeletePost)
				post.GET("/getPosts", postApi.GetPosts)
				post.GET("/get/:id", postApi.GetPost)
			}
			comment := protected.Group("/comment")
			{
				comment.PUT("/create", commentApi.CreateComment)
				comment.GET("/get/:postId", commentApi.GetComment)
			}

		}
	}

	// 启动服务器
	port := ":" + cfg.Server.Port
	logger.Infof("%s %s starting on port %s", cfg.App.Name, cfg.App.Version, cfg.Server.Port)
	logger.Infof("Server running on http://localhost%s", port)

	if err := router.Run(port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
