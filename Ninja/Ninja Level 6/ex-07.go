package main

import "fmt"

func main()  {
	anonymous := func() {
		fmt.Println("Anonymous Func.")
	}
	anonymous()
	fmt.Println("Hello, World!")
}
