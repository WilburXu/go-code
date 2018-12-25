package main

import "fmt"

func main() {
	var list = []int{14, 7, 2, 29, 7, 18, 6, 1, 8, 19, 10}
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

func quickSort(values []int, left int, right int) {
	if left < right {
		// base number
		temp := values[left]

		// set sentinel
		i, j := left, right

		for {
			// 右向左找，找到第一个比分水岭小的数
			for values[j] >= temp && j > i {
				j--
			}

			// 左向右找，找到第一个比分水岭大的数
			for values[i] <= temp && i < j {
				i++
			}

			// 哨兵相遇，break
			if i >= j {
				break
			}

			values[i], values[j] = values[j], values[i]
		}

		// 将分水岭移到哨兵相遇点
		values[left] = values[i]
		values[i] = temp

		// 递归，左右两侧分别排序
		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}
