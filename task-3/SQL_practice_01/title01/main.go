package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

func openDB() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	// 开启db连接
	db := openDB()

	// 1.创建students表
	//createTable(my_db)

	// 2.插入数据
	//studentAdd(my_db)

	// 3.查询数据
	//studentSelect(my_db)

	// 4.更新数据
	//studentUpdate(my_db)

	// 5.删除数据
	studentDelete(db)
}
