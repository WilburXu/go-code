package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "a.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "hello WilburXu\n\r"

	writer := bufio.NewWriter(file)
	for i := 0; i <= 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
