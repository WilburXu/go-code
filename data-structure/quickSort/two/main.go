package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := getTestData(10000000)
	var startTime int64 = time.Now().UnixNano()/1e6;
	quickSort(list, 0, len(list)-1)
	var endTime int64 = time.Now().UnixNano()/1e6;
	fmt.Println("---------------------")
	var divider = endTime - startTime;
	fmt.Println(startTime,endTime);
	fmt.Println(divider)
	//fmt.Println(list)
}


func getTestData(num int32) []int32 {
	var newArray []int32;
	maxRan := num * 10;
	var i int32;
	for i=0;i<num;i++{
		newArray = append(newArray,rand.Int31n(maxRan));
	}

	return newArray;
}

func quickSort(values []int32, left int, right int) {
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
