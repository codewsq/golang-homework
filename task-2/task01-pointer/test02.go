package main

import "fmt"

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func sliceInt(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] * 2
	}
	return numbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sliceInt(numbers)
	fmt.Println(result)
}
