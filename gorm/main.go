package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
)

var (
	ormClient *gorm.DB
)

func init() {
	sourceURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		"root",
		"root",
		"127.0.0.1:13306",
		"live_op")

	var err error
	ormClient, err = gorm.Open("mysql", sourceURL)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	stgs := []int{}
	ormClient.Table("tbl_strategy_user").Where("status = ? and strategy_id > 0", 1).Pluck("strategy_id", &stgs)
	log.Printf("%+v", stgs)
}