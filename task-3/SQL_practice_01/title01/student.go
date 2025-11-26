package main

import (
	"fmt"
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

type Student struct {
	ID    int `gorm:"primaryKey;autoIncrement;column:id"`
	Name  string
	Age   int
	Grade string
}

func createTable(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
func studentAdd(db *gorm.DB) int {
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&student)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("插入成功，插入条数", result.RowsAffected)
	return student.ID
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func studentSelect(db *gorm.DB) []Student {
	var students []Student
	result := db.Where("age > ?", 18).Order("age desc").Find(&students)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("查询成功，查询到的数据条数：", result.RowsAffected, "条，查询数据如下：")
	for _, student := range students {
		fmt.Println(student)
	}
	return students
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func studentUpdate(db *gorm.DB) {
	var students []Student
	result := db.Where("name = ?", "张三").Find(&students)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, student := range students {
		// 方式一
		//my_db.Model(&student).Where("id = ?", student.ID).Update("grade", "四年级")
		// 方式二：
		student.Grade = "四年级"
		db.Save(&student)
	}
	fmt.Print("更新成功")
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func studentDelete(db *gorm.DB) {
	// 方式一
	//result := my_db.Where("age < ?", 15).Delete(&Student{})
	// 方式二
	result := db.Delete(&Student{}, "age < ?", 15)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("删除成功，删除条数：", result.RowsAffected)
}
