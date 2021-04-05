package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const soh = string(1)

type LogSt struct {
	RequestId       string
	TestRespBody    string
	ReleaseRespBody string
	Path            string
}

func main() {
	for i := 0; i <= 9; i++{
		testName := fmt.Sprintf("./log/test/sort/log_%d.txt", i)
		releaseName := fmt.Sprintf("./log/release/sort/log_%d.txt", i)
		respName := fmt.Sprintf("./log/release/resp/log_%d.txt", i)
		test(testName, releaseName, respName)
	}
}

func test(testInputFileName string, releaseInputFileName string, respFileName string) {
	var (
		err error
		respFile *os.File
	)

	testInputFile, _ := os.Open(testInputFileName)
	testInputReader := bufio.NewReader(testInputFile)

	releaseInputFile, _ := os.Open(releaseInputFileName)
	releaseInputReader := bufio.NewReader(releaseInputFile)

	if checkFileIsExist(respFileName) {
		respFile, err = os.OpenFile(respFileName, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
	} else {
		respFile, err = os.Create(respFileName)
	}
	if err != nil {
		panic(err)
	}


	var (
		testReadBool     = true
		testLine         []byte
		testLineSlice    []string
		testRequestId    string
		testRespBody     string
		testPath         string
		testLineCnt      int64
		releaseReadBool  = true
		releaseLine      []byte
		releaseLineSlice []string
		releaseRequestId string
		releaseRespBody  string
		releasePath      string
		releaseLineCnt   int64
	)
	for {
		if testReadBool {
			testLineCnt++
			testLine, _, err := testInputReader.ReadLine()
			if err == io.EOF { // 读到文件末尾就退出
				log.Println("is over")
				break
			}

			if err != nil {
				log.Println(err)
				continue
			}
			testLineSlice = strings.Split(string(testLine), soh)
			testRequestId = testLineSlice[0]
			testRespBody = testLineSlice[1]
			testPath = testLineSlice[2]
			testReadBool = false
		}

		if releaseReadBool {
			releaseLineCnt++
			releaseLine, _, err := releaseInputReader.ReadLine()
			if err == io.EOF { // 读到文件末尾就退出
				log.Println("is over")
				break
			}

			if err != nil {
				log.Println(err)
				continue
			}
			releaseLineSlice = strings.Split(string(releaseLine), soh)
			releaseRequestId = releaseLineSlice[0]
			releaseRespBody = releaseLineSlice[1]
			releasePath = releaseLineSlice[2]
			releaseReadBool = false
		}

		log.Printf("test cnt %d, release cnt %d, test reqId %s, release reqId %s \n", testLineCnt, releaseLineCnt, testRequestId, releaseRequestId)

		if testRequestId < releaseRequestId {
			releaseReadBool = true
			if _, err = releaseInputFile.Seek(int64(len(releaseLine)), 1); err != nil {
				log.Println(err)
				continue
			}
		} else if testRequestId > releaseRequestId {
			testReadBool = true
			if _, err = testInputFile.Seek(int64(len(testLine)), 1); err != nil {
				log.Println(err)
				continue
			}
		} else {
			testReadBool = true
			releaseReadBool = true

			var respRet bool
			if testRespBody == releaseRespBody {
				respRet = true
			}
			sortStr := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%t\n", testRequestId, soh, testRespBody, soh, releaseRespBody, soh, testPath, soh, releasePath, soh, respRet)

			// 创建新的 Writer 对象
			w := bufio.NewWriter(respFile)
			_, err = w.WriteString(sortStr)
			if err != nil {
				log.Println(err)
			}
			_ = w.Flush()

			if _, err = releaseInputFile.Seek(int64(len(releaseLine)), 1); err != nil {
				log.Println(err)
				continue
			}

			if _, err = testInputFile.Seek(int64(len(testLine)), 1); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}


func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}