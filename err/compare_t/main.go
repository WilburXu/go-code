package main

import (
	"errors"
	"log"
)

var ErrNotFound = errors.New("aa")

func main() {
	tOne := errors.New("test")
	tTwo := errors.New("test")

	if tOne == tTwo {
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
