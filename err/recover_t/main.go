package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")

	Go(func() {
		fmt.Println("begin")

		panic("handle error")

		fmt.Println("end")
	})

	time.Sleep(2 * time.Second)
}

func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 捕获，返回500之类的。
				fmt.Println("this is panic")
			}
		}()
		x()
	}()
}