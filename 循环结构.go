package main

import "fmt"

func main() {
	// var i int
	// sum := 0
	// 和C语言相似
	// for i = 1; i < 10; i++ {
	// 	sum = sum + i

	// }
	// fmt.Println(sum)

	// 不写初始值
	// for i < 10 {
	// 	sum += i
	// 	i++
	// }
	// fmt.Println(sum)

	// 所有信息都不写，相当于死循环，要在内部添加循环条件
	// for {
	// 	sum += i
	// 	i++
	// 	if i > 10 {
	// 		break
	// 	}
	// }
	// fmt.Println(sum)

	// range 关键字,相当于遍历
	// for i, s := range "hello" {
	// 	fmt.Printf("%d,%c\n", i, s)
	// }

	// for 循环嵌套,输出100以内的素数
	// var i, j int
	// for i = 0; i < 100; i++ {
	// 	for j = 2; j <= (i / j); j++ {
	// 		if i%j == 0 {
	// 			break
	// 		}
	// 	}
	// 	if j > (i / j) {
	// 		fmt.Printf("%d 是素数\n", i)
	// 	}
	// }

	// 输出九九乘法表
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d*%d=%d  ", j, i, i*j)
	// 	}
	// 	fmt.Print("\n")
	// }

	// 循环控制语句：break、continue、goto
	// break 跳出循环，并执行循环之后的语句，在case语句后跳出执行
	// 以下实例有多重循环，演示了使用标记和不使用标记的区别：
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
	// continue跳出当前循环，执行下一次循环，也可以设置label来指定想continue的循环

	// goto 无条件的转移到指定的行
	goto L1
	fmt.Println("123")
L1:
	fmt.Println("跳转")

	//for循环配合goto打印九九乘法表
	for i := 1; i < 10; i++ {
		j := 1
	LOOP:
		if j <= i {
			fmt.Printf("%dx%d=%d  ", j, i, i*j)
			j++
			goto LOOP
		}
		fmt.Print("\n")
	}

	//for循环配合goto打印100以内素数
	C := 1
L2:
	for C < 100 {
		C++
		for c := 2; c < C; c++ {
			if C%c == 0 {
				goto L2
			}
		}
		fmt.Printf("%d 是素数\n", C)

	}
}
