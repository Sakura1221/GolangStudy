package main

import "fmt"

// 如果类名首字母大写，表示其他包也可以访问
type Hero struct {
	// 如果属性名首字母大写，表示该属性对外能访问，否则只能类的内部访问
	Name string
	Ad int
	Level int
}

/* 值传递方法
// 结构体方法，this表示调用对象，Hero是其类型
func (this Hero) Show() {
	//fmt.Println("Name = ", this.Name)
	//fmt.Println("Ad = ", this.Ad)
	//fmt.Println("Level = ", this.Level)
	fmt.Println("hero = ", this)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	// 当前this是调用该方法对象的拷贝，不能改变原来的值
	this.Name = newName
}
*/

// 引用传递方法
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	// 当前this是调用该方法对象的拷贝，不能改变原来的值
	this.Name = newName
}

func main() {
	// 定义一个对象
	hero := Hero{Name: "ZhangSan", Ad: 100, Level: 1}

	hero.Show()

	hero.GetName()
	hero.SetName("LiSi")

	hero.Show()

}