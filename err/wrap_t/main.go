package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	err := test()
	if err != nil {
		fmt.Printf("%+v", errors.Cause(err))
		fmt.Printf("tt \n %+v", err)
	}
}

func test() (err error){
	fmt.Println("test")

	_, err = os.Open("test.test")
	if err != nil {
		return errors.Wrap(err, "open error")
	}

	return nil
}

