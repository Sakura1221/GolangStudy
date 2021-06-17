package main

import (
	"fmt"
	"time"
)

// 子goroutine，若主goroutine退出，子goroutine也退出
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	// 创建一个go程 去执行newTask() 流程
	go newTask()

	i := 0

	// 死循环
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}