package main

import "fmt"

func main() {
	// 一般的程序都是按照顺序结构执行的，
	a := 10
	// b := 'a'
	// c := a + int(b)
	// fmt.Println(a, b, c)
	// // c表示原始字符，不会将字符转化为ascii编码
	// fmt.Printf("%d,%c,%v", a, b, c)

	// 条件语句，通过判断是否满足条件来决定是否执行指定语句,条件语句有以下几种
	// if语句，布尔表达式
	// if a > 20 {
	// 	fmt.Println("a大于20")
	// } else {
	// 	fmt.Println("a小于20")
	// }

	// 多分支条件语句
	// if a == 10 {
	// 	fmt.Println("true")
	// } else if a > 10 {
	// 	fmt.Println("too big")
	// } else {
	// 	fmt.Println("too small")
	// }

	// if 嵌套语句
	// if a > 10 {
	// 	if b == 'a' {
	// 		fmt.Println("正确")
	// 	} else {
	// 		fmt.Println("错误")
	// 	}
	// } else {
	// 	if a == 10 {
	// 		fmt.Println("a=10")
	// 	} else {
	// 		fmt.Println("a!=10")
	// 	}
	// }

	// switch选择语句，基于不同的条件执行不同的动作，每一个case都是唯一的，默认自带break语句，匹配成功后就不会执行后面的语句，如果我们需要执行后面的 case，可以使用 fallthrough
	// case 后的数据的数据类型必须和变量一致，否则会出错
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	// fallthrough 强制执行后面的case一句，且不会判断是否满足条件
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	// switch 支持多条件匹配
	switch a {
	case 1, 2, 3, 4:
		fmt.Println("1234")
	default:
		fmt.Println("test")
	}

}
