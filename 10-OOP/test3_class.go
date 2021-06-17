package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// =============================

type SuperMan struct {
	Human // SuperMan类继承了Human类的方法
	level int // 新增属性
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"ZhangSan", "male"}

	h.Eat()
	h.Walk()

	// 实例化对象
	// 方式一：定义一个子类对象，先调用父类构造函数，再补充构造子类对象
	// s := SuperMan{Human{"LiSi", "male"}, 88}
	var s SuperMan
	s.name = "LiSi"
	s.sex = "male"
	s.level = 88

	s.Walk() // 父类的方法
	s.Eat() // 子类的方法
	s.Fly() // 子类的方法

	s.Print()
}