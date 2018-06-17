package main

import "fmt"


func main() {
	foo()
	bar()
}

func foo() {
	for i:=1; i<=20; i++ {
	fmt.Println("Foo:", i)
	}
}


func bar() {
	for i:=1; i<=20; i++ {
	fmt.Println("bar:", i)
	}
}
