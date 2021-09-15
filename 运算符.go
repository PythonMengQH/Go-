package main

import "fmt"

/*算术运算符、关系运算符、逻辑运算符、位运算符、赋值运算符、其他运算符*/
func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b // 整数相除结果为整数
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++ // 自增
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--    // 自减
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	fmt.Print(a == b)
	// 逻辑运算符只能对布尔值进行运算
	fmt.Print(true && false)
	fmt.Print(true || false)
	fmt.Print(!false)

	// 位运算符和python中的类似
	// 赋值运算符也和python类似，但是没有幂运算

	// 其他运算符，&表示返回变量的存储地址，*表示指针变量，可以取值
	var age int = 18
	var age1 *int
	age1 = &age
	fmt.Print(&age, &age1)
	fmt.Print(*age1)
}
