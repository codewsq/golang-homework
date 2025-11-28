package database

import (
	"github.com/codewsq/blog/server/config"
	"github.com/codewsq/blog/server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

// ConnectDB 连接数据库
func ConnectDB() error {
	// 获取全局配置实例对象
	cfg := config.GetConfig()
	if cfg == nil {
		log.Fatal("Config not loaded")
	}
	// 获取数据库连接字符串
	dsn := cfg.Database.GetDSN()

	// 配置GORM日志级别
	var gormLogger logger.Interface
	if cfg.Server.Mode == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Second)

	// 自动迁移表结构
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		return err
	}

	DB = db
	log.Println("Database connected successfully")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
