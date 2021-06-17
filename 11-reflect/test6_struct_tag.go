package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() // 取出所有元素

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%s\n", t.Field(i).Name)
		// 解析info
		taginfo := t.Field(i).Tag.Get("info")
		fmt.Println("info: ", taginfo)
		// 解析doc
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("doc: ", tagdoc)
	}
}

func main() {
	var re resume
	// reflect需要传引用
	findTag(&re)
}