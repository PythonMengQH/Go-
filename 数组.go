package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*数组是具有相同唯一类型的一组已编号且长度固定的数据项序列
  数组元素相当于bumpy中的数组，可以通过下标索引来读取或修改*/
func main() {
	// 声明并赋值数组
	// var lst [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(lst)
	// 通过字面量在声明数组的同时快速初始化数组(自动推导数据类型)
	// balance := [5]float64{1.0, 2.3, 4.6, 5.2, 7.9}
	// fmt.Println(balance)

	// 数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自动推导数组长度
	// var lst1 = [...]string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	// fmt.Println(lst1)

	// 如果设置了数组的长度，我们还可以通过指定下标来初始化元素(数组类型默认为0)
	// lst2 := [5]int{1: 2, 4: 99}
	// fmt.Println(lst2)

	// 如果忽略[]中的数字，Go语言会根据元素的个数来设置数组的大小，也叫切片
	// lst3 := []int{1, 2, 3, 4}
	// fmt.Println(lst3)

	// 通过下标索引访问数组元素，和列表取值类似
	// fmt.Println(lst3[2])

	// var n [10]int
	// var i int

	// for i = 0; i < 10; i++ {
	// 	n[i] = i + 100
	// }

	// for _, num := range n {
	// 	fmt.Println(num)
	// }

	// 多维数组声明(这里的多维数组和numpy中的多维数组类似)
	// var twodim [5][4]int
	// 赋值,通过大括号来赋值
	// twodim = [5][4]int{{0, 1, 2, 3}, {1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}, {4, 5, 6, 7}}
	// fmt.Println(twodim)

	// // 通过append方法，为二维数组添加数据
	// values := [][]int{}

	// row1 := []int{1, 2, 3}
	// row2 := []int{4, 5, 6}
	// values = append(values, row1)
	// values = append(values, row2)
	// fmt.Println(values)
	// // 单独输出一行数据
	// fmt.Println(values[0])
	// // 输出单个元素
	// fmt.Println(values[0][0])

	// 访问二维数组，通过行索引和列索引访问元素
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	// var i, j int

	/* 输出数组元素 */
	// for i = 0; i < 5; i++ {
	// 	for j = 0; j < 2; j++ {
	// 		fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
	// 	}
	// }

	// 通过range函数对数组操作得到的结果是每一行的数据
	for _, numb := range a {
		fmt.Println(numb)
	}

	// animals := [][]string{}

	// // 创建三一维数组，各数组长度不同
	// row1 := []string{"fish", "shark", "eel"}
	// row2 := []string{"bird"}
	// row3 := []string{"lizard", "salamander"}

	// // 使用 append() 函数将一维数组添加到二维数组中
	// animals = append(animals, row1)
	// animals = append(animals, row2)
	// animals = append(animals, row3)

	// fmt.Println(animals)

	// 向函数传递数组，需要在定义函数时，声明形参为数组

	var lst = [5]float32{1000.0, 2.0, 3.0, 17.0, 50.0}
	var avg float32
	avg = getAverage(lst, 5)
	fmt.Println("平均值为", avg)

	// 因为浮点数运算有一定的偏差，所以可以通过设置为整数来运算
	d := 1690      // 表示1.69
	b := 1700      // 表示1.70
	c := d * b     // 结果应该是2873000表示 2.873
	fmt.Println(c) // 内部编码
	fmt.Println(float64(c) / 1000000)
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(10))

}

// 在做运算的时候，数据类型需要保持一致
func getAverage(arr [5]float32, size int) float32 {
	var i int
	var avg, sum float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = sum / float32(size)
	return avg
}
