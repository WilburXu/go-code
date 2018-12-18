package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手机Start..")
}

func (p Phone) Stop() {
	fmt.Println("手机Stop")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("Camera Start..")
}

func (c Camera) Stop() {
	fmt.Println("Camera Stop..")
}


func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"huawei"}
	usbArr[1] = Phone{"mini"}
	usbArr[2] = Camera{"suoni"}

	fmt.Println(usbArr)
}
