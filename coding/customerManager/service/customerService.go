package service

import "go-code/coding/customerManager/model"

// 完成对Customer对的数据操作
type CustomerService struct {
	customers []model.Customer	// get customer.struct
	customerNum int
}

