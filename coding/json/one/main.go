package main

import (
	"fmt"
	"wkadmin/json"
)

type Monster struct {
	Name string
	Age int
}

func main() {
	structFunc()
	mapFunc()
	sliceFunc()
}

func structFunc() {
	monster := Monster{
		Name: "WilburXu",
		Age : 25,
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func mapFunc() {
	var data map[string]interface{}
	//使用map,需要make
	data = make(map[string]interface{})

	data["name"] = "WilburXu"
	data["age"] = 25

	ret, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	fmt.Println(string(ret))
}

func sliceFunc() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前，需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "WilburXu"
	m2["age"] = "25"
	m2["address"] = [2]string{"墨西哥","夏威夷"}
	slice = append(slice, m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Println(string(data))
}