package main

import (
	"fmt"
	"math"
)

const s string = "contant";

//Go 支持字符、字符串、布尔和数值 常量 。
func main() {
	fmt.Println(s)

	const n = 500000000
	//常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)
	//数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	fmt.Println(int64(d))
	//当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。举个例子，这里的 math.Sin函数需要一个 float64 的参数。
	fmt.Println(math.Sin(n))

}