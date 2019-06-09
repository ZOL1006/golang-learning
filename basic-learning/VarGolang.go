package main

import "fmt"

var ( // 全局变量
	aa int = 4
	bb bool = true
	// cc := 1 不可出现在函数体外
)
func main() {
	var a string = "String"
	fmt.Println(a)
	var b, c = 1, 2
	fmt.Println(b - 1 + c + aa)
	newvar :=1 // := 必须为新的声明 变量 //只能在函数体内出现
	fmt.Println(newvar)
	say()
}

func say()  {
	fmt.Println(bb)
}