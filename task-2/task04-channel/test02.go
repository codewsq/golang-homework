package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制
*/

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 确保函数退出时通知WaitGroup
	defer close(ch) // 关键：生产完成后关闭通道
	for i := 0; i < 100; i++ {
		ch <- i + 1
		fmt.Println("生产者发送数据：", i+1)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("生产者完成")
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 确保函数退出时通知WaitGroup
	for value := range ch {
		fmt.Println("消费者接收数据：", value)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("消费者完成")
}

func main() {
	c := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(c, &wg)
	go consumer(c, &wg)

	wg.Wait()
	fmt.Println("main end...")
}
