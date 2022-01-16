package main

import (
	"errors"
	"fmt"
)

func main() {
	//_, err := os.Open("/test.txt")
	//if err != nil {
	//	fmt.Println("open failed, err:", err)
	//	return
	//}

	//errOne := fmt.Errorf("fmt.Errorf() err, http status is %d", 404)
	//fmt.Printf("errOne errType is %T，err is %v\n", errOne, errOne)


	if ok, err := err.(VipErr); ok {

	}
	wErrOne := errors.New("this is one ")
	wErrTwo := fmt.Errorf("this is two %w", wErrOne)
	fmt.Printf("wErrOne type is %T err is %v \n", wErrOne, wErrOne)
	fmt.Printf("wErrTwo type is %T err is %v \n", wErrTwo, wErrTwo)
}
