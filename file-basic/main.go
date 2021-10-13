package main

import (
	"log"
	"os"
)

func main() {
	createFile()

}

// 创建文件
// 重复操作会覆盖文件
// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
func createFile() {
	newFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	defer newFile.Close()
}


