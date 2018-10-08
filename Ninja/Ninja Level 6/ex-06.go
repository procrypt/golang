package main

import "fmt"

func main()  {
	func() {
		fmt.Println("Anonymous Func.")
	}()
	fmt.Println("Hello, World!")
}
