package main

import "fmt"

//可变入参函数
func main() {
	//常规调用
	sum(1, 2)
	sum(1, 2, 3)

	//如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
	nums := []int{1, 2, 3, 4}
	sum(nums...)


	//多余的参数会报错
	//numArr := [5]int{6,7,8,9,10}
	//数组不可以
	//sum(numArr...)
}


//这个函数使用任意数目的 int 作为参数。
//类似于java的动态参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
