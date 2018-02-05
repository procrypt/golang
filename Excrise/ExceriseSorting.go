package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	//sort.SliceStable(studyGroup, func(i, j int) bool {
	//	return studyGroup[i] < studyGroup[j]
	//})
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
