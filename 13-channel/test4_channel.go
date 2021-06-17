package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 通过close关闭channel（管道内数据清空后关闭）
		close(c)
	}()

/*	// 主go程死循环，如果管道不close，会发生死锁
	for {
		// ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		// 分号表达式
		if data, ok := <- c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}*/
	// range可以阻塞地从管道迭代取数据
	// 效果与上面代码相同
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished...")
}