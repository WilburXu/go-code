package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccc")

	// 终止协程
	runtime.Goexit()

	fmt.Println("ddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaaa")

		test()

		fmt.Println("bbbbbb")
	}()

	for {
	}
}