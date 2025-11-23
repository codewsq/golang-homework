package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	name string
	fn   func()
}

func goroutine_1(wg *sync.WaitGroup, name string, t func()) {
	defer wg.Done() // 在每个 goroutine 中，我们使用 defer wg.Done() 来确保在 goroutine 退出时调用 Done() 方法，将计数器减 1

	taskStart := time.Now()
	t()
	duration := time.Since(taskStart)

	fmt.Printf("任务名称：%s 完成，耗时: %v\n", name, duration)
}

func main() {
	// 任务列表
	tasks := make([]task, 4)
	tasks[0] = task{"快速任务", func() { time.Sleep(100 * time.Millisecond) }}
	tasks[1] = task{"中等任务", func() { time.Sleep(500 * time.Millisecond) }}
	tasks[2] = task{"慢速任务", func() { time.Sleep(1 * time.Second) }}
	tasks[3] = task{"数据处理", func() { time.Sleep(300 * time.Millisecond) }}

	// 声明并初始化了一个 WaitGroup，此时计数器为 0
	var wg sync.WaitGroup
	start := time.Now()

	fmt.Println("开始并发执行任务...")

	// 并发执行所有任务
	for _, task := range tasks {
		wg.Add(1) // 在循环中，每次启动一个 goroutine 之前，我们调用 wg.Add(1) 来增加 WaitGroup 的计数器，表示我们要等待一个 goroutine 完成。这里我们为每个任务增加 1，所以计数器最终为任务的数量
		go goroutine_1(&wg, task.name, task.fn)
	}

	// 等待所有任务完成:在主 goroutine 中，我们调用 wg.Wait() 来阻塞，直到 WaitGroup 的计数器归零。这意味着所有我们启动的 goroutine 都已经执行完毕，
	// 然后主 goroutine 才会继续执行后面的代码（即打印总执行时间）
	wg.Wait()

	fmt.Printf("所有任务执行完毕！\n")
	fmt.Printf("总执行时间: %v\n", time.Since(start))
}
