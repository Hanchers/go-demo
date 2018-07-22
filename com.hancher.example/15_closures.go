package main

import "fmt"

//Go 支持通过 闭包来使用 匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
// 所有函数外变量全局可用的语言，都应该会有闭包的概念，用来解决全局变量作用域的问题。
func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

//这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。这个返回的函数使用闭包的方式 隐藏 变量 i。
// 闭包最典型的应用就是计数器的实现
// 这是一个自增计数器
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}