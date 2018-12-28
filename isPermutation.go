package main

import (
	"fmt"
	"strings"
)

func isPermutation(s1, s2 string) bool {
	s1 = strings.Replace(s1, " ", "", -1)
	s2 = strings.Replace(s2, " ", "", -1)

	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if strings.Contains(s2, string(s1[i])) {
			s2 = strings.Replace(s2, string(s1[i]), "",1)
		}
	}
	return len(s2) == 0
}

func main() {
	fmt.Println(isPermutation("driving", "drivign"))
}