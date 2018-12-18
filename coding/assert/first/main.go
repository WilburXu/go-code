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
	fmt.Println("phone start..")
}

func (p Phone) Stop() {
	fmt.Println("phone stop..")
}

func (p Phone) Call() {
	fmt.Println("phone is calling..")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("camera is start..")
}

func (c Camera) Stop() {
	fmt.Println("camera is stop..")
}

type Computer struct {

}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"huawei"}
	usbArr[1] = Phone{"mini"}
	usbArr[2] = Camera{"kane"}

	var computer Computer
	for _, v := range usbArr {
		fmt.Println(v)
		computer.Working(v)
		fmt.Println("--")
	}
}