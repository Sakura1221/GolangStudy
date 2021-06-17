package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 运行结束")
		fmt.Println("goroutine 正在运行...")

		// 向管道中放数据
		// 管道默认是无缓冲的，如果管道内有数据，那么子go程会阻塞
		c <- 666 // 将666发送给c
	}()

	// 从管道中取数据
	// 内含同步关系，保证子go程先向管道内放数据，否则就阻塞
	num := <- c // 从c中接受数据，并赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}