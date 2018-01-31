package main

import "fmt"

type foo int

func main() {

	var myage foo
	myage = 44
	fmt.Printf("%T, %v\n", myage, myage)

	var yourage int
	yourage = 29
	fmt.Printf("%T, %v\n", yourage, yourage)

	//fmt.Println(myage + yourage)

	fmt.Println(int(myage) + yourage)

}
