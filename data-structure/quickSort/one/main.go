package main

import "fmt"

func main() {
	var list = []int{7, 2, 29, 18, 6, 1, 8, 19, 10}
	ret := quickSort(list)
	fmt.Println(ret)
}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	mid := list[0]
	var (
		leftArr []int
		rightArr []int
	)

	for i := 1; i < len(list); i++ {
		if list[i] > mid {
			rightArr = append(rightArr, list[i])
		} else {
			leftArr = append(leftArr, list[i])
		}
	}

	fmt.Println(leftArr)
	fmt.Println(rightArr)

	leftArr = quickSort(leftArr)
	rightArr = quickSort(rightArr)

	leftArr = append(leftArr, mid)
	leftArr = append(leftArr, rightArr...)

	return leftArr
}


