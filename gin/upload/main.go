package main

import (
	"fmt"
	"time"
)

type A struct {
	Name string `json:"name"`
	Name1 string `json:"name1"`
}

func main() {
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	fmt.Println(endTime.Unix()-time.Now().Unix())
	fmt.Println(endTime.Format("2006/01/02 15:04:05"))
	//router := gin.Default()
	//router.GET("/upload", func(c *gin.Context) {
	//
	//	ret := A{
	//		Name: "ed2k://|file|%C6%AF%C1%C1%20%C4%A3%CC%D8%20%D3%EB%C5%D6%20%B4%E5%B9%C3%20%CD%E6SM.avi|119424364|673D9EBC82D1A32F0AFC4F35B053BA5D|/",
	//		Name1: "ed2k://|file|Æ¯ÁÁ Ä£ÌØ ÓëÅÖ ´å¹Ã ÍæSM.avi|119424364|673D9EBC82D1A32F0AFC4F35B053BA5D|/",
	//	}
	//
	//	retJson, _ := json.Marshal(ret)
	//
	//	c.JSON(200, string(retJson))
	//
	//})
	//
	//router.Run(":9090")
}
