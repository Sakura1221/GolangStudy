package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// tty: pair<type:*os.File, "/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// r:pair<tyep: , value:>
	var r io.Reader
	// r:pair<type:*os.File, "/dev/tty"文件描述符>
	r = tty

	// r:pair<
	var w io.Writer
	// 将r强制转换成Writer，小括号是强制类型转换
	// w:pair<type:*os.File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer)

	// Write函数参数是byte数组
	w.Write([]byte("HELLO THIS is A TEST!!!\n"))
}