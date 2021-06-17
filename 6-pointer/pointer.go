package main

import "fmt"

/* 值传递
func swap(a int, b int){
	var tmp int
	tmp = a
	a = b
	b = tmp
}
*/

// 引用传递
func swap(pa *int, pb *int){
	var tmp int
	tmp = *pa // tmp = main::a
	*pa = *pb // main::a = main::b
	*pb = tmp // main::b = tmp
}

func main(){
	var a int = 10
	var b int = 20

	// swap
	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int // 二级指针
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}