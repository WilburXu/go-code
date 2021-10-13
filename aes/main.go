package main

import (
	"encoding/json"
	"go-code/aes/aespkcs5"
	"log"
)

type GetMarginParam struct {
	Gcid    string `json:"gcid"`
	Version string `json:"version"`
}

var key = "???"


func main() {
	var getMarginParam = GetMarginParam{
		Gcid : "111",
		Version: "0.0.4",
	}


	jsonStr, _ := json.Marshal(getMarginParam)
	encodeStr, _ := aespkcs5.Encrypt(key, string(jsonStr))
	log.Println(encodeStr)


	ret, err := aespkcs5.Decrypt(key, string(encodeStr))
	log.Println(string(ret))
	log.Println(err)

	decodeInfo()
}

func decodeInfo() {
	str := "HyYhr1hzqxvGNAynqG20rF+kR9kXEbFmedN2eOJ7JkmjniVq0QBfiLIwYMw+94LvU2fvvtQ+R1AHuf0DBenivA=="
	ret, err := aespkcs5.Decrypt(key, str)
	log.Println(string(ret))
	log.Println(err)
}