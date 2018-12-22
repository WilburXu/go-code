package main

import (
	"fmt"
	"go-code/coding/customerManager/service"
	"go-code/coding/customerManager/model"
)

type customerView struct {
	key  string
	loop bool
	customerService *service.CustomerService
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
				this.add()
			case "2":
				fmt.Println("修 改 客 户")
			case "3":
				this.delete()
			case "4":
				this.list()
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

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

func (this *customerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(this.customerService.CustomerId + 1, name, gender, age, phone, email)
	this.customerService.Add(customer)
}

func (this *customerView) delete() {
	fmt.Println("delete id?")
	id := -1
	fmt.Scanln(&id)

	if id == -1 {
		return
	}

	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if (this.customerService.Delete(id)) {
			fmt.Println("success")
		} else {
			fmt.Println("faild")
		}
	}
}

func main() {
	customerView := customerView{
		key: "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}