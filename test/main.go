package main

import (
	"bufio"
	"log"
	"os"
)

const soh = string(1)

type LogSt struct {
	RequestId       string
	TestRespBody    string
	ReleaseRespBody string
	Path            string
}

func main() {
	testInputFile, _ := os.Open("./log/test/sort/log_0.txt")
	//testInputFile, _ := os.Open("./access.log")
	testInputReader := bufio.NewReader(testInputFile)

	//var testLine []byte
	//var testErr error

	releaseInputFile, _ := os.Open("./log/release/sort/log_0.txt")
	releaseInputReader := bufio.NewReader(releaseInputFile)

	//releasePos := 0
	for i := 0; ; i++ {
		_, _, err := testInputReader.ReadLine()
		if err != nil {
			log.Println(err)
			return
		}

		//log.Println(string(testLine))

		//testLineSlice := strings.Split(string(testLine), soh)
		//requestId := testLineSlice[0]
		//TestRespBody := testLineSlice[1]
		//Path := testLineSlice[2]
		//ReleaseRespBody := lineSlice[2]

		for {
			releaseLine, _, err := releaseInputReader.ReadLine()
			if err != nil {
				log.Println(err)
				return
			}

			if _, err := releaseInputReader.Discard(len(releaseLine) + 1); err != nil {
				log.Println(err)
				return
			}

			log.Println(releaseLine)
		}

		//break
	}

	//_, err = testInputFile.Seek(64, 0)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//line, _, err = testInputReader.ReadLine()
	//log.Println(string(line))

	//fPos := 0 // or saved position
	//for i := 1; ; i++ {
	//	testLine, testErr = testInputReader.ReadBytes('\n')
	//	log.Printf("[line:%d pos:%d] %q\n", i, fPos, testLine)
	//
	//	if testErr != nil {
	//		break
	//	}
	//	fPos += len(testLine)
	//}

	//file, _ := os.Open("./access.log")
	//defer file.Close()
	//// 偏离位置，可以是正数也可以是负数
	//var offset int64 = 5
	//
	//// 用来计算offset的初始位置
	//// 0 = 文件开始位置
	//// 1 = 当前位置
	//// 2 = 文件结尾处
	//var whence int = 0
	//newPosition, err := file.Seek(offset, whence)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Just moved to 5:", newPosition)
	//// 从当前位置回退两个字节
	//newPosition, err = file.Seek(-2, 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Just moved back two:", newPosition)
	//// 使用下面的技巧得到当前的位置
	//currentPosition, err := file.Seek(0, 1)
	//fmt.Println("Current position:", currentPosition)
	//// 转到文件开始处
	//newPosition, err = file.Seek(0, 0)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Position after seeking 0,0:", newPosition)
}
