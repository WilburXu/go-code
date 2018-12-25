package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:/go-code/README.md")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(file)

	err = file.Close()

	if err != nil {
		fmt.Println(err)
	}
}
