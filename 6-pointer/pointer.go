package main

import "fmt"


func swap(a int ,b int) {
	fmt.Println("a和b内部交换前：a = ", a, " b = ", b)
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Println("a和b内部交换后：a = ", a, " b = ", b)
}


func swap2(pa *int, pb *int) {
	// 定义一个中间变量
	var temp int
	// 交换
	temp = *pa //temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}


func main() {
	var a int = 10
	var b int = 20

	fmt.Println("a和b交换前：a = ", a, " b = ", b)

	fmt.Println("a和b交换前的地址：a = ", &a, " b = ", &b)

	swap(a, b)

	fmt.Println("a和b交换后：a = ", a, " b = ", b)
	fmt.Println("a和b交换后的地址：a = ", &a, " b = ", &b)

	// 定义一个指针变量 p
	var p *int

	// 指针变量的存储地址到p
	p = &a

	// 打印p的地址
	fmt.Println(&a)
	// 打印p的值
	fmt.Println(p)

	// 定义一个二级指针变量 pp
	var pp **int

	pp = &p

	// 打印pp的地址
	fmt.Println(&p)
	// 打印pp的值
	fmt.Println(pp)
}