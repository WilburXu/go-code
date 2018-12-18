package main

import "fmt"

type Student struct {
	name string
	age int
}


func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	stu := Student{}

	TypeJudge(n1, n2, n3, name ,address, n4, stu, &stu)
}


func TypeJudge(items... interface{}) {
	for index, x := range items {
		switch x.(type) {
			case bool :
				fmt.Printf("the %v is bool val is %v \n", index, x)
			case float32 :
				fmt.Printf("the %v is float32 val is %v \n", index, x)
			case float64 :
				fmt.Printf("the %v is float64 val is %v \n", index, x)
			case int, int32, int64 :
				fmt.Printf("the %v is int val is %v \n", index, x)
			case string :
				fmt.Printf("the %v is string val is %v \n", index, x)
			case Student :
				fmt.Printf("the %v is Student val is %v \n", index, x)
			case *Student :
				fmt.Printf("the %v is *Student val is %v \n", index, x)
			default :
				fmt.Printf("the %v is unknown val is %v \n", index, x)
		}
	}
}