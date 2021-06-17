package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
			case c <- x:
				// 如果c可写，则该case就会进来
				x, y = y, x + y
			case <-quit: // 注意这里检测的不是管道内的数据，而是管道是否可读
				fmt.Println("quit")
				return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		// 通过该管道发送退出信号
		quit <- 0
	}()

	// main go
	fibonacii(c, quit)
}