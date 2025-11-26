package main

import (
	"fmt"
	"github.com/codewsq/gorm/title02/blog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
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
	user, err := blog.SelectInfoByUserId2(openDB(), 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询指定用户下的所有文章和评论信息：")
	fmt.Println("用户名称：", user.Name)
	posts := user.Posts
	for _, post := range posts {
		fmt.Println("----发表的文章 - 标题：", post.Title, ", 内容：", post.Content, "，评论条数：", len(post.Comments))
		comments := post.Comments
		for _, comment := range comments {
			fmt.Println("========文章评论：", comment.Context)
		}
		fmt.Println("******************************************************************")
	}

	fmt.Println("\n============================查询评论最多的文章信息=====================================\n")
	post, err := blog.SelectPostByMaxCom(openDB())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("评论数最多的文章 - 标题：", post.Title, ", 内容：", post.Content)
}
