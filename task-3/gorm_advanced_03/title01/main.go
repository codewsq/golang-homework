package main

import (
	"fmt"
	"github.com/codewsq/gorm/title01/blog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
func openDB() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	err := blog.CreateTable(openDB())
	if err != nil {
		panic(err)
	}
	fmt.Println("Gorm创建模型数据库表成功")

	errUser := blog.SaveInfo(openDB())
	if err != nil {
		panic(errUser)
		return
	}
	fmt.Println("保存信息成功")
}
