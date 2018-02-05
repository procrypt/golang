package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{
		"Zeno",
		"John",
		"Al",
		"Jenny",
	}

	fmt.Println(s)
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}
