package main

import "fmt"
/*
	四种变量的声明方式
*/

// 前三种可以用来声明全局变量
var gA int
var gB int = 100
var gC = 200

// 方法四:=只能在函数体内声明变量


func main() {
	// 方法一：声明一个变量，默认的值是0
	var a int
	fmt.Println("a =", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)

	// 方法三：初始化时，省略数据类型，通过值自动匹配当前变量的数据类型
	var c= 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	// 方法四：省去var关键字，直接自动匹配
	d := 100
	fmt.Println("d =", d)
	fmt.Printf("type of d = %T\n", d)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, ", yy =", yy)
	var kk, ll = 100, "Sakura"
	fmt.Println("kk =", kk, ", ll =", ll)

	// 多行的多变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv =", vv, ", jj =", jj)
}