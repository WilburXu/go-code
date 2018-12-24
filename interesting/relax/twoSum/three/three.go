package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	var target int = 13

	ret := twoSum(nums, target)
	fmt.Println(ret)
}

/**
	- 一遍哈希方法
	- 时间复杂度：O(n)O(n)， 我们只遍历了包含有 nn 个元素的列表一次。在表中进行的每次查找只花费 O(1)O(1) 的时间。
	- 空间复杂度：O(n)O(n)， 所需的额外空间取决于哈希表中存储的元素数量，该表最多需要存储 nn 个元素。
 */

func twoSum(nums []int, target int) []int {
	var vMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		index, ok := vMap[complement]
		if ok {
			return []int{i, index}
		}

		vMap[nums[i]] = i
	}

	return nil
}
