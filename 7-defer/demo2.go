/*
知识点二：defer和return的先后顺序
return 比 defer 先执行
*/

package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc() // defer要写在return之前
	return returnFunc()
}

func main() {
	returnAndDefer()
}