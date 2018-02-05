package main

import "fmt"

type animal struct{
	speak string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func spec(a ...interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"Woof"}, true}
	firi := cat{animal{"Meow"}, false}

	spec(fido, firi)
	spec(firi)
}
