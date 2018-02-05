package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{7, 8, 5, 4, 6, 2, 1, 3, 9}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
