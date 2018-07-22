package main

import "fmt"
// go 中的值处理与Java类似
func main() {
	fmt.Println("go"+"lang");
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("7.0/3 =", 7.0/3)
	fmt.Println("7/3.0 =", 7/3.0)
	fmt.Println("7/3 =", 7/3)
	fmt.Println("7%3 =", 7%3)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
