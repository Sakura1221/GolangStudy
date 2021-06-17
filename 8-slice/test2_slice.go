package main

import "fmt"

// 引用传递
func printArray_(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4, 5} // 动态数组，切片 slice
	fmt.Printf("myArray type is %T\n", myArray)

	printArray_(myArray)
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}