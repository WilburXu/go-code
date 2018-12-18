package main

import "fmt"

type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "自带爬树功能..")
}

type LittleMoney struct {
	Monkey
}

func (this *LittleMoney) Flying() {
	fmt.Println(this.Name, "学习后，会飞了..")
}

func main() {
	monkey := LittleMoney{}
	monkey.Name = "悟空"

	monkey.climbing()
	monkey.Flying()
}