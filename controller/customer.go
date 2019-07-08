package controller
import (
	"../model"
)

//该结构体完成对customer的操作
//增删改查
type CustomerController struct {
	Customers []model.Customer
	customerNum int
}

func NewCc() *CustomerController {
	//初始化一个客户
	customerController := &CustomerController{}
	customerController.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",20,"110","zs@qq.com")
	customerController.Customers = append(customerController.Customers,customer)
	return customerController
}

//返回客户切片
func (cc *CustomerController) List() []model.Customer {
	return cc.Customers
}

//添加客户到切片
func (cc *CustomerController) Add(customer model.Customer) bool {
	cc.customerNum++
	customer.Id = cc.customerNum
	cc.Customers = append(cc.Customers,customer)
	return true
}

//根据id查找客户在切片中对应的下标，没有返回-1
func (cc *CustomerController) FindById(id int) int {
	index := -1
	for i := 0; i < len(cc.Customers); i++ {
		if cc.Customers[i].Id == id {
			index = i
		}
	}
	return index
}

//根据id从切片删除

func (cc *CustomerController) Delete(id int) bool {
	index := cc.FindById(id)
	if index == -1 {
		return false
	}
	// 删除操作
	cc.Customers = append(cc.Customers[0:index],cc.Customers[index+1:]...)
	return true
}


func (cc *CustomerController) Update(customer model.Customer){
	index := cc.FindById(customer.Id)
	if(customer.Name != "") {
		cc.Customers[index].Name = customer.Name
	}
	if(customer.Gender != "") {
		cc.Customers[index].Gender = customer.Gender
	}
	if(customer.Age != 0) {
		cc.Customers[index].Age = customer.Age
	}
	if(customer.Phone != "") {
		cc.Customers[index].Phone = customer.Phone
	}
	if(customer.Email != "") {
		cc.Customers[index].Email = customer.Email
	}
}
