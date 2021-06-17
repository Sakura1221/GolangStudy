package main

import "fmt"

// 声明一种新的数据类型myint，是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	author string
}

func changeBook(book Book) {
	// 值传递
	book.author = "666"
}

func changeBook2(book *Book) {
	// 引用传递
	book.author = "666"
}

func main() {
/*	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a) // main.myint  */

	var book1 Book
	book1.title = "Golang"
	book1.author = "ZhangSan"

	fmt.Printf("%v\n", book1)
	changeBook(book1) // 值传递
	fmt.Printf("%v\n", book1)
	changeBook2(&book1) // 引用传递
	fmt.Printf("%v\n", book1)


}