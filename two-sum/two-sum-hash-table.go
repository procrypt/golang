package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		compliment := target - nums[i]
		if _, ok := m[compliment]; ok {
			return []int{i, m[compliment]}
		} else {
			m[v] = i
		}
	}
	return nil
}

func main() {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(a)
}
