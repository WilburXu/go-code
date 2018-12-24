package main

import "fmt"

func main() {
	var list = []int{7, 2, 29, 18, 6, 1, 8, 19, 10}
	quickSort(list)
}

func quickSort(list []int) {

	mid := list[0]
	var (
		leftArr []int
		rightArr []int
	)

	fmt.Println(mid)

	for i := 0; i <= len(list); i++ {
		if list[i] > mid {
			rightArr[] = list[i]
		} else {
			
		}
	}
}


