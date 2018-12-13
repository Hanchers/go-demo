package main

import "fmt"

//变量
func main() {

	//显式声明一个变量
	var a  = 'a' ;
	//此处会报错： cannot use 111 (type int) as type string in assignment
	// a = 111
	fmt.Println(a)
	//fmt.Println(s)

	//你可以申明一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)
	//Go 将自动推断已经初始化的变量类型。
	var d = true
	fmt.Println(d)
	//声明变量且没有给出对应的初始值时，变量将会初始化为零值 。例如，一个 int 的零值是 0。
	var e int
	fmt.Println(e)
	//:= 语句是申明并初始化变量的简写，例如这个例子中的 var f string = "short"。
	f := "short"
	fmt.Println(f)

}
