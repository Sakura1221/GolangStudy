package main

import (
	"fmt"
	"time"
)

func main() {
	// 缓冲为3的channel
	c := make(chan int, 3)

	fmt.Println("len(c) =", len(c), ", cap(c) =", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送的元素=", i, ", len(c)=", len(c), ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i ++ {
		num := <- c
		fmt.Println("num = ", num)
	}

	fmt.Println("主go程结束")
}