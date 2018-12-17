package main

import (
	"container/list"
	"fmt"
	"time"
)

/**
	golang中 slice 与 list的性能测试
	insert:
		slis 创建时间： 945.5144ms
		lists 创建时间： 14.2619099s

		slis 创建时间： 887.3863ms
		lists 创建时间： 14.7121065s


	select/find:
		slis 遍历时间： 48.0768ms
		lists 遍历时间： 332.886ms

		slis 遍历时间： 47.1245ms
		lists 遍历时间： 388.0044ms
 */

func main() {
	//create()
	//find()
}

func create() () {
	var maxNum int = 100000000

	sliTime := time.Now()
	slis := make([]int, 10)
	for i := 0; i < maxNum; i++ {
		slis = append(slis, i)
	}
	fmt.Println("slis 创建时间：", time.Now().Sub(sliTime).String())

	listTime := time.Now()
	lists := list.New()
	for j := 0; j < maxNum; j++ {
		lists.PushBack(j)
	}
	fmt.Println("lists 创建时间：", time.Now().Sub(listTime).String())
}

func find() {
	var maxNum int = 100000000

	slis := make([]int, 10)
	lists := list.New()

	for i := 0; i < maxNum; i++ {
		slis = append(slis, i)
		lists.PushBack(i)
	}

	var tmp int
	sliTime := time.Now()
	for _, v := range slis {
		tmp = v
	}
	fmt.Println("slis 遍历时间：", time.Now().Sub(sliTime).String())

	var lTmp interface{}
	listTime := time.Now()
	for e := lists.Front(); e != nil; e = e.Next() {
		lTmp = e.Value
	}
	fmt.Println("lists 遍历时间：", time.Now().Sub(listTime).String())

	fmt.Println(tmp, lTmp)
}
