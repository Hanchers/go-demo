package main

import "fmt"


func plus (a int,b int) int  {
	return a+b
}

type person struct {
	name string
}

func hello(user string) {
	fmt.Println(user+" say： hello world!")
}

func (p person) hello(user string){
	fmt.Println(user+" say:hello "+p.name)
}

type SayHello interface {
	hello(name string)
}


func main() {
	// function
	res := plus(1,2)
	fmt.Println("1+2 =", res)

	// function and method
	hello("function")
	p := person{"Hancher"}
	p.hello("method")

	//interface
	toSayHello(p,"method")

	//chan
	msg := make(chan string)
	go func() {msg<-"chan"}()
	user:= <-msg
	hello(user)
	p.hello(user)
}

func toSayHello(sayHello SayHello,user string)  {
	sayHello.hello(user)
}