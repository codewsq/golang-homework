package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func sendNumber(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(c)
	for i := 1; i <= 10; i++ {
		c <- i
	}
}

func receiveNumber(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range c {
		fmt.Println("打印接收的数据：", value)
	}
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go sendNumber(c, &wg)

	go receiveNumber(c, &wg)

	wg.Wait()
	fmt.Println("main end")

}
