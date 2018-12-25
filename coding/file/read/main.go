package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	filePath := "C:/go-code/README.md"
	// 一次性全读
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))


	// 缓冲区 一次读一部分
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("over")
}
