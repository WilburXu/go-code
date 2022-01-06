package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

)

func main() {
	err := mid()
	if err != nil {
		log.Printf("%+v", errors.Cause(err))
		log.Printf("tt \n %+v", err)
	}
}

func mid() (err error) {
	return test()
}

func test() (err error){
	fmt.Println("test")

	_, err = os.Open("test.test")
	if err != nil {
		return errors.Wrap(err, "open error")
	}

	return nil
}

