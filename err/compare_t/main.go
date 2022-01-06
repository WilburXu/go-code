package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNotFound = errors.New("aa")

func main() {
	t1 := errors.New("a")
	t2 := errors.New("a")

	if t1 == t2 {
		log.Println("ok")
	} else {
		log.Println("no")
	}


	err := fmt.Errorf("access denied: %w", ErrNotFound)
	if errors.Is(err, ErrNotFound) {
		log.Println("is ")
	} else {
		log.Println("no is")
	}
}
