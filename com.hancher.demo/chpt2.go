package main

import "fmt"


func plus (a int,b int) int  {
	return a+b
}

type person struct {
	name string
}
//function
func hello(user string) {
	fmt.Println(user+" say： hello world!")
}

//method
func (p person) hello(user string)  {
	s:=user+" say:hello "+p.name
	fmt.Println(s)
	//return s
}
func  (p person) hello2(user string) string  {
	s:=user+" say:hello "+p.name
	fmt.Println(s)
	return s
}

type SayHello interface {
	hello(name string)
	//hello2(a string) string
}
func toSayHello(sayHello SayHello,user string)  {
	sayHello.hello(user)
}

func main() {
	// function
	res := plus(1,2)
	fmt.Println("1+2 =", res)

	// function and method
	hello("function")
	p := person{"Hancher"}
	p.hello("method1")

	//interface
	toSayHello(p,"method2")

	//chan
	//var msg1 chan string;
	//msg1=make(chan  string)
	msg:= make(chan string)
	go func() {msg<-"chan"}()
	user:= <-msg
	hello(user)
	p.hello(user)
}

