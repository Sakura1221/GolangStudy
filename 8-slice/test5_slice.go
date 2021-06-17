package main

import "fmt"

func main() {
	s := []int{1, 2, 3} // len = 3, cap = 3

	// [0, 1)
	s1 := s[0:1]
	fmt.Println(s1)

	// 浅拷贝，修改s1也会修改s
	s1[0] = 100
	fmt.Println(s)

	// 深拷贝copy 可以讲底层数组的slice一起进行拷贝
	s2 := make([]int, 3)
	// 将s中的值拷贝到s2中
	copy(s2, s)
}