package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素1，numbers len = 4, [0, 0, 0, 1], numbers cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	// 向numbers切片追加一个元素2，numbers len = 5, [0, 0, 0, 1, 2], numbers cap = 5
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 容量已满，倍增容量
	numbers = append(numbers, 3)
	// 向numbers切片追加一个元素3，numbers len = 6, [0, 0, 0, 1, 2, 3], numbers cap = 10
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 容量与初始化的长度相同
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	// 继续append，cap也会倍增，len = 4，cap = 6
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}