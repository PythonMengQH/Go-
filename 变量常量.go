package main

import (
	"fmt"
	"unsafe"
)

// 这种不带声明格式的变量只能在函数体内出现
// f := 1
// func main() {
// 	// 单个变量赋值
// 	var a int
// 	a = 123
// 	// 多个变量进行赋值
// 	var b, c string
// 	b = "456"
// 	// 根据值自行判断变量类型
// 	var d = 3.14
// 	// 声明加赋值
// 	e := "hello"
// 	fmt.Print(a, b, c, d, e)
// }

// Go语言常量，不能被修改的值

// func main() {
// 	// 常量的定义必须要赋值，可以省略type，让其自己推断
// 	const a int = 1
// 	const b = "abd"
// 	// 同时声明多个常量
// 	const c, d, e = 2, 3.14, 'a'
// 	// iota 在声明常量的时候才能使用，每遇到const关键字都会被重置为0，const每增加一行常量声明iota会自增1
// 	const (
// 		f = iota
// 		g = iota
// 	)
// 	fmt.Println(a, b, c, d, e, f, g)
// }

// iota用法
func main() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Print(unsafe.Sizeof(d))
}
