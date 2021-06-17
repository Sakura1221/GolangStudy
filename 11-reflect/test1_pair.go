package main

import "fmt"

func main() {
	var a string
	// pair<statictype:string, value:"sakura">
	a = "sakura"

	// pair<type:string, value:"sakura">
	var allType interface{}
	allType = a // 赋值就传递

	value, ok := allType.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value =", value)
	}
}