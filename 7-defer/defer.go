package main

import "fmt"

/**
这段代码是用 Go 语言编写的一个 main 函数，该函数使用了 defer 关键字。代码的功能是在程序执行完毕前，延迟打印两句话。

下面是代码的解释：

在 main 函数开始时，使用了 defer 关键字，在这个关键字后面跟着的是在函数结束前需要执行的操作。这里有两个 defer 语句，第一个将打印 "main end1"，第二个将打印 "main end2"。由于它们是延迟执行的，所以它们将在 main 函数中的其他操作执行完毕后才执行。

defer 语句下方，有两个 fmt.Println 语句。第一个 fmt.Println 将在控制台打印 "main::hello go 1"，第二个 fmt.Println 将在控制台打印 "main::hello go 2"。这两行语句将在 defer 语句之前执行，因为它们没有被延迟。

main 函数中没有循环或者条件语句来改变程序的流程，因此，fmt.Println("main::hello go 1") 和 fmt.Println("main::hello go 2") 将按顺序执行，接着是 "main end2" 和 "main end1"。这是因为后执行的 defer 语句会覆盖先执行的 defer 语句。

defer 关键字通常用于资源的释放，比如关闭文件描述符，释放锁，或者做一些清理工作。在这个例子中，defer 用于在 main 函数结束前执行一些清理或者收尾性的操作。

*/
func main() {
	// 使用 defer 关键字确保 fmt.Println("main end1") 在函数结束前被执行
	defer fmt.Println("main end1")
	// 使用 defer 关键字确保 fmt.Println("main end2") 在函数结束前被执行
	defer fmt.Println("main end2")

	// 打印字符串 "main::hello go 1"
	fmt.Println("main::hello go 1")
	// 打印字符串 "main::hello go 2"
	fmt.Println("main::hello go 2")
}

