package main

import "fmt"

func main() {

	greeting := make(map[string]string)
	greeting["Abhi"] = "Hi!"
	greeting["Bobby"] = "Hey!"

	fmt.Println(greeting)
}