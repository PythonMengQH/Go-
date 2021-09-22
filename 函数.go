/*函数是基本的代码块，用于执行一个任务，Go语言最少有一个main()函数
函数声明告诉了编译器函数的名称，返回类型和参数*/

package main

import "fmt"

// 函数定义，由关键字func声明，function_name指的是函数名称，后面的是参数列表和返回值及类型
// func function_name([parameter list]) [return_types]{
// 	函数体
// }

//实例
// func max(num1, num2 int) (result int) {
// 	if num1 > num2 {
// 		result = num1
// 	} else {
// 		result = num2
// 	}
// 	return
// }

// 函数返回多个值
// func swap(x, y string) (e, f string) {
// 	e, f = y, x
// 	return
// }

// func main() {
// 	a, b := 100, 200
// 	c, d := "baidu", "google"
// 	// 调用函数，在定义函数时定义的形参需要传入实参
// 	result := max(a, b)
// 	// 返回多个值需要使用多个变量进行接收
// 	e, f := swap(c, d)
// 	fmt.Println(e, f)
// 	fmt.Println(result)

// 匿名函数，如果定义在主函数外，则必须使用变量进行赋值
// var f = func() {
// 	fmt.Println("匿名函数")
// }

// // 引用传递值,不推荐
// func main() {
// 	/* 定义局部变量 */
// 	var a int = 100
// 	var b int = 200

// 	fmt.Printf("交换前，a 的值 : %d\n", a)
// 	fmt.Printf("交换前，b 的值 : %d\n", b)

// 	/* 调用 swap() 函数
// 	&a 指向 a 指针，a 变量的地址
// 	&b 指向 b 指针，b 变量的地址
// 	*/
// 	// fmt.Println(&a, *&a)
// 	swap(&a, &b)

// 	fmt.Printf("交换后，a 的值 : %d\n", a)
// 	fmt.Printf("交换后，b 的值 : %d\n", b)
// 	f()
// 	fmt.Printf("%T", f)
// 	func() {
// 		fmt.Println("匿名函数")
// 	}()
// }

// func swap(x *int, y *int) {
// 	var temp int
// 	temp = *x /* 保存 x 地址上的值 */
// 	*x = *y   /* 将 y 值赋给 x */
// 	*y = temp /* 将 temp 值赋给 y */
// }

// Go语言闭包函数

func getSequence() func() (j int) {
	i := 0
	return func() (j int) {
		i += 1
		j = i
		return
	}
}

func main() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 带参数的匿名函数
	func(a, b int) int {
		sum := a + b
		fmt.Println(sum)
		return sum
	}(20, 10)

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

// 函数中的不定长参数
func sum(num ...int) int {
	sum := 0
	for _, i := range num {
		sum += i
	}
	return sum
}
