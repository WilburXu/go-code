package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("aa")


func main() {
	err := mid()
	err = errors.WithMessage(err, "33333")

	if err != nil {
		log.Printf("%+v", errors.Cause(err))
		log.Printf("tt \n %+v", err)
	}



	if errors.Is(err, ErrNotFound) {
		log.Println("111")
	} else {
		log.Println("222")
	}

	time.Sleep(2 * time.Second)
	log.Println("run done")

}

func mid() (err error) {

	return test()
}

func test() (err error) {
	fmt.Println("test")

	_, err = os.Open("test.test")
	if err != nil {
		//err := fmt.Errorf("access denied: %w", ErrNotFound)
		return errors.Wrap(err, "open error")
	}

	return nil
}
