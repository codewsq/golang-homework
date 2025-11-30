package database

import (
	"github.com/codewsq/blog/server/config"
	mylogger "github.com/codewsq/blog/server/logger"
	"github.com/codewsq/blog/server/models"
	"github.com/sirupsen/logrus"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB 连接数据库
func ConnectDB() error {
	cfg := config.GetConfig()
	if cfg == nil {
		mylogger.Fatal("Config not loaded")
	}

	dsn := cfg.Database.GetDSN()

	// 记录连接信息（不包含密码）
	mylogger.WithFields(logrus.Fields{
		"host":   cfg.Database.Host,
		"port":   cfg.Database.Port,
		"dbname": cfg.Database.DBName,
	}).Info("Connecting to database")

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
		mylogger.WithFields(logrus.Fields{
			"error": err.Error(),
			"dsn":   dsn, // 生产环境不要记录完整DSN
		}).Error("Failed to connect to database")
		return err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		mylogger.Errorf("Failed to get database instance: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Second)

	// 自动迁移表结构
	mylogger.Info("Starting database migration")
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		mylogger.Errorf("Failed to migrate database: %v", err)
		return err
	}

	DB = db
	mylogger.Info("Database connected successfully")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
