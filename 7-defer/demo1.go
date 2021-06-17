/*
知识点一：defer的执行顺序
基于栈，先进后出
*/

package main

import "fmt"

func func1(){
	fmt.Println("A")
}

func func2(){
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main(){
	if (false) {
		// 写入defer关键字，在函数结束前触发，类似java中的final
		// 用栈存储，先进后出
		defer fmt.Println("main end1")
		defer fmt.Println("main end2")

		fmt.Println("main: hello go 1")
		fmt.Println("main: hello go 2")
	} else { // 输出C B A
		defer func1()
		defer func2()
		defer func3()
	}
}