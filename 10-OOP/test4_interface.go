package main

import "fmt"

// 定义一个interface类型的类，作为统一接口
// interface本质是一个指针，子类将内部方法全部重载，就可以指向子类
type AnimalIF interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string // 获取动物的种类
}

// 具体的类
type Cat struct {
	color string // 猫的颜色
}

// 实现接口中的方法，相当于继承了interface
func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string // 猫的颜色
}

// 实现接口中的方法，相当于继承了interface
func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// 函数参数类型设置成接口，就可以接收多种类型，实现多态
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	var animal AnimalIF // 接口的数据类型，父类指针
	animal = &Cat{"Green"}
	animal.Sleep() // 调用的就是Cat的Sleep()方法

	animal = &Dog{"Yellow"}
	animal.Sleep() // 调用的就是Dog的Sleep()方法

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)
}