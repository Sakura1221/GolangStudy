package main

import "fmt"

// 固定长度数组传参，长度不同的数组，是不同类型
// 值拷贝
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
}

func main() {
	// 固定长度的数组，默认初始化为0
	var myArray1 [10]int
	// 初始化前4个为1,2,3,4
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	//for i := 0; i < 10; i ++ {}
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// 查看数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
}