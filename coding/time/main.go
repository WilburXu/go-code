package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("current xtime", time.Now())
	//fmt.Println("format time:", time.Now().Format("2006-01-02 15:04"))
	//
	//// 字符串转 Time
	//withNanos := "2006-01-02 15:04:05"
	//t, _ := time.Parse(withNanos, "2013-10-05 18:30:50")
	//fmt.Println(t)

	SteadyTime()
	DiffTime()
}


func DiffTime() {
	//var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	local, _ := time.LoadLocation("Local")
	endTimeString := "2019-03-12 00:30:00"
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endTimeString, local)

	//endTime, _ := time.Parse("2006-01-02 15:04:05", endTimeString)
	fmt.Println(endTime)
	fmt.Println(time.Now())
	duration := time.Now().Sub(endTime)
	fmt.Println(duration.Hours())
}

// 定时每天固定时间
func SteadyTime() {
	//timeNow := time.Now() //获取当前时间
	//zeroHour := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
	//next := zeroHour.Add(- time.Hour * 48)
	//
	//fmt.Println(zeroHour)
	//fmt.Println(next)
}

