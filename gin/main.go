package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
)

var (
	nonce = "9I0fMA(GC#qG"
)

type LogSt struct {
	RequestId string
	RespBody  string
	Path      string
}

type LogListSt []LogSt

const soh = string(1)

func main() {
	//ReadTestFileContentOfLine("./log/cleaning.2.log")
	//ReadTestFileSort("./log/test/split/")

	//ReadReleaseFileContentOfLine("./log/team.proxy.access.1617307201.log")
	//ReadReleaseFileContentOfLine("./log/team.proxy.access.1617314401.log")
	//ReadReleaseFileContentOfLine("./log/team.proxy.access.1617321601.log")
	//ReadReleaseFileContentOfLine("./log/team.proxy.access.1617328801.log")
	//ReadReleaseFileSort("./log/release/split/")
}


func ReadReleaseFileSort(dirPath string) {
	files, _ := ioutil.ReadDir(dirPath)
	for _, file := range files {
		logList := make(LogListSt, 0)

		if file.IsDir() {
			continue
		}

		if fileExt := path.Ext(file.Name()); fileExt != ".txt" {
			continue
		}

		i := 0
		log.Println(file.Name(), " begin")
		splitFilePath := fmt.Sprintf("%s%s", dirPath, file.Name())
		r, _ := os.Open(splitFilePath)
		s := bufio.NewScanner(r)
		for s.Scan() {
			i++
			log.Printf("%s:%d", file.Name(), i)
			line := s.Text()
			lineSlice := strings.Split(line, soh)
			if len(lineSlice) < 3 {
				continue
			}

			logData := LogSt{
				RequestId: lineSlice[0],
				RespBody:  lineSlice[1],
				Path:      lineSlice[2],
			}
			logList = append(logList, logData)
		}

		sort.Sort(logList)

		var (
			sortFile *os.File
			err      error
		)

		sortFilePath := fmt.Sprintf("./log/release/sort/%s", file.Name())
		if checkFileIsExist(sortFilePath) { //如果文件存在
			sortFile, err = os.OpenFile(sortFilePath, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		} else {
			sortFile, err = os.Create(sortFilePath) //创建文件
		}
		if err != nil {
			log.Println(err)
		}

		for _, v := range logList {
			sortStr := fmt.Sprintf("%s%s%s%s%s\n", v.RequestId, soh, v.RespBody, soh, v.Path)

			// 创建新的 Writer 对象
			w := bufio.NewWriter(sortFile)
			_, err = w.WriteString(sortStr)
			if err != nil {
				log.Println(err)
			}
			_ = w.Flush()
		}
	}
}

// 读取测试k8s集群的access log
// 按照 request_id 的第一个字符的ascii求余，写入新文件
// 新文件按照 request_id 排序
func ReadReleaseFileContentOfLine(filename string) {
	splitFileObjs := make(map[int]*os.File, 0)

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		log.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	log.Println("file size=", size)

	buf := bufio.NewReader(file)
	i := 0
	for {
		i++
		log.Println(i)
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				log.Println("File read ok!")
				break
			} else {
				log.Println("Read file error!", err)
				return
			}
		}
		lineSlice := strings.Split(line, soh)
		if len(lineSlice) < 4 {
			log.Printf("i=%d lineSlice < 4 \n", i)
			continue
		}

		u, err := url.Parse(lineSlice[3])
		if err != nil {
			log.Println(err)
			continue
		}

		requestId := strings.Trim(lineSlice[0], "\"")
		respBody := strings.Trim(lineSlice[1], "\"")
		urlPath := strings.Trim(u.Path, "/")

		modNum := int(math.Mod(float64(requestId[0]), 10))
		modFileName := fmt.Sprintf("./log/release/split/log_%d.txt", modNum)
		var splitFile *os.File
		if v, ok := splitFileObjs[modNum]; ok {
			splitFile = v
		} else {
			if checkFileIsExist(modFileName) { //如果文件存在
				splitFile, err = os.OpenFile(modFileName, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
			} else {
				splitFile, err = os.Create(modFileName) //创建文件
			}
			if err != nil {
				log.Println(err)
			}

			splitFileObjs[modNum] = splitFile
		}

		splitStr := fmt.Sprintf("%s%s%s%s%s\n", requestId, soh, respBody, soh, urlPath)
		w := bufio.NewWriter(splitFile) //创建新的 Writer 对象
		_, err = w.WriteString(splitStr)
		if err != nil {
			log.Println(err)
		}
		_ = w.Flush()
	}
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 获取此 slice 的长度
func (p LogListSt) Len() int { return len(p) }

// 根据元素的requestId降序排序 （此处按照自己的业务逻辑写）
func (p LogListSt) Less(i, j int) bool {
	return p[i].RequestId > p[j].RequestId
}

// 交换数据
func (p LogListSt) Swap(i, j int) { p[i], p[j] = p[j], p[i] }


func ReadTestFileSort(dirPath string) {
	files, _ := ioutil.ReadDir(dirPath)
	for _, file := range files {
		logList := make(LogListSt, 0)

		if file.IsDir() {
			continue
		}

		if fileExt := path.Ext(file.Name()); fileExt != ".txt" {
			continue
		}

		i := 0
		log.Println(file.Name(), " begin")
		splitFilePath := fmt.Sprintf("%s%s", dirPath, file.Name())
		r, _ := os.Open(splitFilePath)
		s := bufio.NewScanner(r)
		for s.Scan() {
			i++
			log.Printf("%s:%d", file.Name(), i)
			line := s.Text()
			lineSlice := strings.Split(line, soh)
			if len(lineSlice) < 3 {
				continue
			}

			logData := LogSt{
				RequestId: lineSlice[0],
				RespBody:  lineSlice[1],
				Path:      lineSlice[2],
			}
			logList = append(logList, logData)
		}

		sort.Sort(logList)

		var (
			sortFile *os.File
			err      error
		)

		sortFilePath := fmt.Sprintf("./log/test/sort/%s", file.Name())
		if checkFileIsExist(sortFilePath) { //如果文件存在
			sortFile, err = os.OpenFile(sortFilePath, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		} else {
			sortFile, err = os.Create(sortFilePath) //创建文件
		}
		if err != nil {
			log.Println(err)
		}

		for _, v := range logList {
			sortStr := fmt.Sprintf("%s%s%s%s%s\n", v.RequestId, soh, v.RespBody, soh, v.Path)

			// 创建新的 Writer 对象
			w := bufio.NewWriter(sortFile)
			_, err = w.WriteString(sortStr)
			if err != nil {
				log.Println(err)
			}
			_ = w.Flush()
		}
	}
}

// 读取测试k8s集群的access log
// 按照 request_id 的第一个字符的ascii求余，写入新文件
// 新文件按照 request_id 排序
func ReadTestFileContentOfLine(filename string) {
	splitFileObjs := make(map[int]*os.File, 0)

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		log.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	log.Println("file size=", size)

	buf := bufio.NewReader(file)
	i := 0
	for {
		i++
		log.Println(i)
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				log.Println("File read ok!")
				break
			} else {
				log.Println("Read file error!", err)
				return
			}
		}
		lineSlice := strings.Split(line, soh)
		if len(lineSlice) < 4 {
			log.Printf("i=%d lineSlice < 3 \n", i)
			continue
		}

		urlSlice := strings.Split(lineSlice[3], " ")
		u, err := url.Parse(urlSlice[1])
		if err != nil {
			log.Println(err)
			continue
		}

		requestId := strings.Trim(lineSlice[0], "\"")
		respBody := strings.Trim(lineSlice[1], "\"")
		urlPath := strings.Trim(u.Path, "/")

		modNum := int(math.Mod(float64(requestId[0]), 10))
		modFileName := fmt.Sprintf("./log/test/split/log_%d.txt", modNum)
		var splitFile *os.File
		if v, ok := splitFileObjs[modNum]; ok {
			splitFile = v
		} else {
			if checkFileIsExist(modFileName) { //如果文件存在
				splitFile, err = os.OpenFile(modFileName, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
			} else {
				splitFile, err = os.Create(modFileName) //创建文件
			}
			if err != nil {
				log.Println(err)
			}

			splitFileObjs[modNum] = splitFile
		}

		splitStr := fmt.Sprintf("%s%s%s%s%s\n", requestId, soh, respBody, soh, urlPath)
		w := bufio.NewWriter(splitFile) //创建新的 Writer 对象
		_, err = w.WriteString(splitStr)
		if err != nil {
			log.Println(err)
		}
		_ = w.Flush()
	}
}
