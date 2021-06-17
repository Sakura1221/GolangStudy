package main

import (
	_ "GolangStudy/5-init/lib1"      // 匿名导入，不能使用方法（会执行init函数）
	mylib2 "GolangStudy/5-init/lib2" // 给包起别名
	. "GolangStudy/5-init/lib3" // 将包中方法都加入到当前包（不建议，避免接口冲突）
)

func main(){
	// lib1.Lib1Test()

	// lib2.Lib2Test()
	mylib2.Lib2Test()
	Lib3Test()
}