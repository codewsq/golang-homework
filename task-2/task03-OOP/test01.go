package main

import (
	"fmt"
	"reflect"
)

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格
*/

// 定义一个 Shape 接口
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

// 实现 Shape 接口
func (r *Rectangle) Area() {
	fmt.Println("Rectangle.Area...")
}
func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle.Perimeter...")
}

type Circle struct {
}

// 实现 Shape 接口
func (c *Circle) Area() {
	fmt.Println("Circle.Area...")
}
func (c *Circle) Perimeter() {
	fmt.Println("Circle.Perimeter...")
}

func main() {
	var shape Shape
	shape = new(Rectangle)
	shape.Area()
	shape.Perimeter()
	fmt.Println("Shape is ", reflect.TypeOf(shape))
	fmt.Println("---------------------------")
	shape = &Circle{}
	shape.Area()
	shape.Perimeter()
	fmt.Println("Shape is ", reflect.TypeOf(shape))
}
