package main

import "fmt"

//指针，学过C的都知道
//指针的使用会是程序更灵活，但也令变量的作用域变相扩大，且不易追踪。
//自由与安全永远是一对矛盾体
func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	//通过 &i 语法来取得 i 的内存地址，例如一个变量i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	//指针也是可以被打印的。
	fmt.Println("pointer:", &i)
}

//zeroval 有一个 int 型参数，所以使用值传递。
//zeroval 将从调用它的那个函数中得到一个 ival形参的拷贝。
func zeroval(ival int) {
	ival = 0
}
//zeroptr 有一和上面不同的 *int 参数，意味着它用了一个 int指针。
//函数体内的 *iptr 接着解引用 这个指针，从它内存地址得到这个地址对应的当前值。对一个解引用的指针赋值将会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}
