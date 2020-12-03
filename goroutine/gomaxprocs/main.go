package main

import (
	"runtime"
	"fmt"
)

func main(){
	// 指定核数
	n := runtime.GOMAXPROCS(1)
	fmt.Println(n)

	for {
		go fmt.Println(1)

		fmt.Println(0)
	}
}

