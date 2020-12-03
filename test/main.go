package main

import (
	"log"
	"time"
)

func main() {
	log.Println("start")
	go test()

	log.Println("end")
}

func test() {
	for {
		log.Println("aa")
		time.Sleep(2 * time.Second)
	}
}
