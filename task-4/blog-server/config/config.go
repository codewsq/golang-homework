package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

// Config 全局配置结构体
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	App      AppConfig      `mapstructure:"app"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	Charset         string `mapstructure:"charset"`
	ParseTime       bool   `mapstructure:"parse_time"`
	Loc             string `mapstructure:"loc"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret          string `mapstructure:"secret"`
	ExpirationHours int    `mapstructure:"expiration_hours"`
}

// AppConfig 应用配置
type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

// GlobalConfig 全局配置实例
var GlobalConfig *Config

// LoadConfig 加载配置文件
func LoadConfig(configPath string) error {
	// 如果未指定配置文件路径，使用默认路径
	if configPath == "" {
		// 先检查当前目录
		if _, err := os.Stat("config.yaml"); err == nil {
			configPath = "config.yaml"
		} else if _, err := os.Stat("config/config.yaml"); err == nil {
			configPath = "config/config.yaml"
		} else {
			log.Fatal("Config file not found")
		}
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// 读取环境变量覆盖配置
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 解析配置到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}

	log.Printf("Config loaded successfully from: %s", configPath)
	return nil
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return GlobalConfig
}

// GetDSN 获取数据库连接字符串
func (d *DatabaseConfig) GetDSN() string {
	return d.Username + ":" + d.Password + "@tcp(" + d.Host + ":" + strconv.Itoa(d.Port) + ")/" + d.DBName + "?charset=" + d.Charset + "&parseTime=" + strconv.FormatBool(d.ParseTime) + "&loc=" + d.Loc
}
