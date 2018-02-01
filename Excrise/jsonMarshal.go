package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int
}

func main() {
	p1 := Person{"Abhishek", "Singh", 25}
	byteSlice, _ := json.Marshal(p1)
	fmt.Printf("%T\n", byteSlice)
	fmt.Printf(string(byteSlice))

}