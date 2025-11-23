package main

import "fmt"

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
type Person struct {
	// 包含 Name 和 Age 字段
	Name string
	Age  int
}

type Employee struct {
	Person            // 组合 Person 结构体
	EmployeeID string // 添加 EmployeeID 字段
}

// 实现一个 PrintInfo() 方法,输出员工的信息。
func (e *Employee) PrintInfo() {
	fmt.Println("EmployeeID:", e.EmployeeID, ",Name:", e.Name, ",Age:", e.Age)
}

func main() {
	emp := Employee{Person{"James", 20}, "0001"}

	emp.PrintInfo()
}
