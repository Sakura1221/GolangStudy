package main

import "fmt"

// 传引用
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"
	changeValue(cityMap)

	fmt.Println("---------")

	// 遍历
	printMap(cityMap)
}