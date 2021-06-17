package main

import (
	"fmt"
	"time"
)

func main() {

/*  // 用go创建承载一个形参为空，返回值为空的一个匿名函数
	go func() {
		defer fmt.Println("A.defer")

		// 内嵌一个子函数，也是匿名函数
		func() {
			defer fmt.Println("B.defer")
			// return // 这里的return只能退出子函数
			runtime.Goexit() // 这个方法可以退出整个go程
			fmt.Println("B")
		}() // 最后要加括号，进行调用，不加括号只是定义

		fmt.Println("A")
	}()*/

	// 用go创建承载一个有形参，返回值的一个匿名函数
	go func(a int,b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	// 死循环，无限休眠
	for {
		time.Sleep(1 * time.Second)
	}
}