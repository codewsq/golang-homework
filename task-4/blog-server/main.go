package main

import (
	"github.com/codewsq/blog/server/api"
	"github.com/codewsq/blog/server/config"
	"github.com/codewsq/blog/server/database"
	"github.com/codewsq/blog/server/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	if err := config.LoadConfig(""); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	cfg := config.GetConfig()

	// 设置Gin运行模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化数据库连接
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 创建Gin路由
	router := gin.Default()

	// 添加基础中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 初始化控制器
	authApi := &api.AuthApi{}
	postApi := &api.PostApi{}
	commentApi := &api.CommentApi{}

	// 健康检查路由
	router.GET("/health", func(c *gin.Context) {
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
	log.Printf("%s %s starting on port %s", cfg.App.Name, cfg.App.Version, cfg.Server.Port)
	log.Printf("Server running on http://localhost%s", port)

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
