package main

import "fmt"

//Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
//返回 a+b
func plus(a int, b int) int {
	return a+b
}

func main() {
	res := plus(1,2)
	fmt.Println("1+2 =", res)
}
