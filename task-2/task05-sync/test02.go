package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func main() {
	var counter int64 // 必须使用 int64 类型，因为 atomic 包需要确定的内存大小
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个协程递增1000次
			for j := 0; j < 1000; j++ {
				// 使用原子操作递增计数器（重点-核心）
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 使用原子操作读取最终值（重点-核心）
	finalValue := atomic.LoadInt64(&counter)
	fmt.Printf("\n最终计数器值: %d (期望值: 10000)\n", finalValue)

	if finalValue == 10000 {
		fmt.Println("测试通过：计数器值正确")
	} else {
		fmt.Println("测试失败：计数器值不正确")
	}
}
