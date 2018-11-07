package two_sum

import "fmt"

func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		for j := k + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{k, j}
			}
		}
	}
	return nil
}

func main() {
	a := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(a)
}
