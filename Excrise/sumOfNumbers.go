package main

import "fmt"

func main() {
	sum := 0
	for i:= 0; i < 50; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum =  sum + i
			}
		}
		fmt.Println(sum)
	}