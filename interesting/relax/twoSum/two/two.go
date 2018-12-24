package main

func main() {
	var nums = []int{2, 7, 11, 15}
	var target int = 9

	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	var vMap = make(map[int]int)
	for key, num := range nums {
		vMap[num] = key
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		index, ok := vMap[complement]
		if ok && index != i {
			return []int{i, index}
		}
	}
	return nil
}