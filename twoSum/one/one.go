package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	var target int

	fmt.Scanln(&target)

	tmp := twoSum(nums, target)
	fmt.Println(tmp)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j<len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}