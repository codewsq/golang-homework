package main

import (
	"fmt"
	"github.com/codewsq/gorm/title03/blog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func openDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
	posts := []blog.Post{
		{Title: "mysql", Content: "mysql安装内容", UserID: 1},
		{Title: "Spring", Content: "MVC内容", UserID: 1},
	}
	err := blog.InsertPost(openDB(), &posts)
	if err != nil {
		fmt.Println("插入文章信息失败")
	}
	fmt.Println("*********************************************************")
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	deleteErr := blog.DeleteComment(openDB(), 3)
	if deleteErr != nil {
		fmt.Println("评论删除失败")
	}
}
