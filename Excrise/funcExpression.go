package main

import "fmt"

func main() {

	function := func(n int) (int, bool) {
		result := n / 2
		if n%2 == 0 {
			return result, true
		} else {
			return result, false
		}
	}
	fmt.Println(function(2))
}
