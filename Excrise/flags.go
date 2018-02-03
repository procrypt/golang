package main

import (
	"flag"
	"time"
	"fmt"
)

var period = flag.Duration("Duration", 1*time.Second, "Sleep Period")

func main() {
	flag.Parse()
	fmt.Println(&period)
	*period = 5
	fmt.Println(*period)
	fmt.Println(&period)
	fmt.Printf("Sleeping for %v", *period)
	time.Sleep(*period)
	fmt.Println()
}