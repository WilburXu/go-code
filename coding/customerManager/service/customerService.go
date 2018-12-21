package service

import "go-code/coding/customerManager/model"

// 完成对Customer对的数据操作
type CustomerService struct {
	Customers   []model.Customer // get customer.struct
	CustomerNum int
	CustomerId  int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.CustomerId = 1
	customerService.CustomerNum = 1
	customerM := model.Customer{
		Id:     customerService.CustomerId,
		Name:   "WilburXu",
		Gender: "boy",
		Age:    25,
		Phone:  "13424300000",
		Email:  "WilburXu@gmail.com",
	}
	customerService.Customers = append(customerService.Customers, customerM)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.Customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.Customers = append(this.Customers, customer)
	return true
}