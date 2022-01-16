package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNotFound = errors.New("not found")

func main() {
	errOne := errors.New("test")
	errTwo := errors.New("test")
	fmt.Printf("错误类型为%T \n", errOne)

	if errOne == errTwo {
		log.Println("Equal")
	} else {
		log.Println("notEqual")
	}

	//err := fmt.Errorf("access denied: %w", ErrNotFound)
	//if errors.Is(err, ErrNotFound) {
	//	log.Println("is ")
	//} else {
	//	log.Println("no is")
	//}
}
