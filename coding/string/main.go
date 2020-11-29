package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello world!")
}

func uintToString() {
	var num uint64 = 17
	strconv.FormatUint(uint64(num), 10)
}

func intToString() {
	var num int = 17
	strconv.Itoa(num)
}
