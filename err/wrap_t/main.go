package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

var ErrNotFound = errors.New("aa")


func main() {
	err := mid()
	err = errors.WithMessagef(err, "%d ", 3333)

	if err != nil {
		log.Printf("cause is %+v \n", errors.Cause(err))
		log.Printf("strace tt %+v \n", err)
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
	_, err = os.Open("test.test")
	if err != nil {
		//err = fmt.Errorf("access denied: %w", ErrNotFound)
		return errors.Wrap(err, "open error")
	}

	return nil
}
