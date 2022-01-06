package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var (
	Wait sync.WaitGroup
	Cnt int64
)

func main() {
	for rt := 1; rt <= 8; rt++ {
		Wait.Add(1)
		go Routine(rt)
		Wait.Wait()
	}

	log.Println("ret", Cnt)
}

func Routine(id int) {
	defer Wait.Done()
	for count := 0; count < 2; count++ {
		value := Cnt
		time.Sleep(time.Nanosecond)
		atomic.AddInt64(&value, 1)
		Cnt = value
		log.Println(id)
	}
}

