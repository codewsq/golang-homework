package my_db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var OpenDB *gorm.DB

func openDBFunc() {
	fmt.Println("未开启db连接，开始创建。。。")
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"
	openDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	OpenDB = openDB

}

func init() {
	openDBFunc()
}
