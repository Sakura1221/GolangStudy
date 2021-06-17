package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {

}
// Book重载了两个接口内的方法
// 类型可以强制转换成任意一种接口
// 但转换过程中pair不变
func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// b: pair<tyep:Book, value:>
	b := &Book{}

	// r: pair<tyep:, value:>
	var r Reader
	// r: pair<tyep:Book, value:book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	// r:pair<type:Book, value:book{}地址>
	w = r.(Writer) // 这里的转换能成功，是因为w r具体的type是一致的
	w.WriteBook()
}