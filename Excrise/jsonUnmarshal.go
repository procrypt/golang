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
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	byteSlice := []byte(`{"First":"Abhishek","Last":"Singh","Age":25}`)
	json.Unmarshal(byteSlice, &p1)

	fmt.Println("------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T", p1)


}

