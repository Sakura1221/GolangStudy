package lib1

import "fmt"

// lib1包提供的API
// 函数名首字母大写，表明是对外开放函数，外部可以调用
func Lib1Test(){
	fmt.Println("lib1 test()...")
}

func init(){
	fmt.Println("lib1 init()...")
}