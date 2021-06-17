/*
动态数组初始化
*/
package main

import "fmt"

func main() {
	// 方法一：声明slice是一个切片，并且初始化，默认值日生1, 2, 3, 长度len是3
	// slice1 := []int{1, 2, 3}

	// 方法二：声明slice1是一个切片，但是并没有给slice分配空间
	// var slice1 []int
	// 调用make开辟类型为[]int的空间，长度为3，默认初始化为0
	// slice1 = make([]int, 3)
	// slice1[0] = 100

	// 方法三：声明并分配空间，默认初始化为0
	// var slice1 []int = make([]int, 3)
	// 可简写成
	slice1 := make([]int, 3)


	// %v直接输出所有值
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 判断一个slice是否为空
	// go语言里nil表示null
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 分配了空间")
	}
}