package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Sakura", 18}
	// reflect需要传引用
	GetFieldAndMethod(user)
}

func GetFieldAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	// 通过type获取里面的字段
	// 1.获取interface的reflect.Type，通过Type得到NumField（字段数量），进行遍历
	// 2.通过Field()方法得到每个field的数据名与类型
	// 3.通过field的Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s(%v) = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取里面的方法，并调用
	// 1.获取interface的reflect.Type，通过Type得到NumMethod（方法数量），进行遍历
	// 2.通过Method()方法得到每个方法名称和类型
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type) // Call: func(main.User)
	}
}