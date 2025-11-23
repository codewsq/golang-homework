package main

import (
	"fmt"
	"sync"
)

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func main() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程递增1000次
			mutex.Lock() // 加锁
			for j := 0; j < 1000; j++ {
				counter++ // 临界区操作
			}
			mutex.Unlock() // 解锁
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	fmt.Printf("最终计数器值: %d (期望值: 10000)\n", counter)

	if counter == 10000 {
		fmt.Println("测试通过：计数器值正确")
	} else {
		fmt.Println("测试失败：计数器值不正确")
	}
}
