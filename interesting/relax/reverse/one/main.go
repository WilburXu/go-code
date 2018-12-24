package main

import "fmt"

func main() {
	x := reverse(1534236469)
	fmt.Println(x)
}

func reverse(n int) int {
	var revNum int32
	for y := int32(n); y != 0; y /= 10 {
		temp := revNum * 10 + y % 10
		if temp/10 != revNum {
			return 0
		}
		revNum = temp
	}
	return int(revNum)
}
