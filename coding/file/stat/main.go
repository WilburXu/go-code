package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main() {
	filePath := "a.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 读到文件末尾就退出
			break
		}

		for _, v := range str {
			switch {
				case v >= 'a' && v <= 'z':
					fallthrough // 穿透
				case v >= 'A' && v <= 'Z':
					count.ChCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				default :
					count.OtherCount++
			}
		}
	}

	//输出统计的结果看看是否正确
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}
