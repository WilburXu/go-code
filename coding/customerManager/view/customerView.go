package main

import "fmt"

type customerView struct {
	key  string
	loop bool
}

func (this *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("\t 1 添 加 客 户")
		fmt.Println("\t 2 修 改 客 户")
		fmt.Println("\t 3 删 除 客 户")
		fmt.Println("\t 4 客 户 列 表")
		fmt.Println("\t 5 退      出")
		fmt.Print("请选择(1-5):")

		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				fmt.Println("添 加 客 户")
			case "2":
				fmt.Println("修 改 客 户")
			case "3":
				fmt.Println("删 除 客 户")
			case "4":
				fmt.Println("客 户 列 表")
			case "5":
				this.loop = false
			default:
				fmt.Println("你的输入有误，请重新输入...")
		}

		if this.loop == false {
			break
		}
	}

	fmt.Println("Your had got out...")
}

func main() {
	customerView := customerView{
		key: "",
		loop: true,
	}

	customerView.mainMenu()
}