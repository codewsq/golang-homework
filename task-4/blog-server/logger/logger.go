package logger

import (
	"github.com/codewsq/blog/server/config"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// InitLogger 初始化日志
func InitLogger() error {
	cfg := config.GetConfig()
	if cfg == nil {
		return nil
	}

	Log = logrus.New()

	// 设置日志级别
	level, err := logrus.ParseLevel(cfg.Log.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)

	// 设置日志格式
	if cfg.Log.Format == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	// 设置日志输出
	var output io.Writer
	if cfg.Log.Output == "file" {
		// 确保日志目录存在
		dir := filepath.Dir(cfg.Log.FilePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		file, err := os.OpenFile(cfg.Log.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		output = file
	} else {
		output = os.Stdout
	}

	Log.SetOutput(output)

	return nil
}

// GetLogger 获取日志实例
func GetLogger() *logrus.Logger {
	return Log
}

// 封装常用日志方法
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

// WithFields 添加结构化字段
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Log.WithFields(fields)
}
