package main

import (
	"fmt"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行
*/
// 打印奇数
func oddNumberPrint() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("goroutine-1 oddNumberPrint 打印奇数：", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// 打印偶数
func evenNumberPrint() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("goroutine-1 evenNumberPrint 打印偶数：", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go oddNumberPrint()

	go evenNumberPrint()

	// 15秒后结束主程
	time.Sleep(15 * time.Second)
	fmt.Println("main goroutine 结束")
}
