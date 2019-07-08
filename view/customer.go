package view
import (
	"fmt"
	"../controller"
	"../model"
)

type customerView struct {
	key string	//接受用户输入
	customerController *controller.CustomerController
}

func NewCv() customerView {
	return customerView {
		key:"",
		customerController:controller.NewCc(),
	}
}

func (cv *customerView) ShowMenu() {
	loop:
	for {
		fmt.Println("----------客户信息管理系统----------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退    出")
		fmt.Print("请选择(1 - 5):")
		
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":break loop
		default:fmt.Println("输入有误..")
		}
	}
	fmt.Println("退出成功...")
}

func (cv *customerView) list() {
	customers := cv.customerController.List()
	fmt.Println("--------------客户列表--------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0;i < len(customers);i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------客户列表完成------------\n")
}

func (cv *customerView) add() {
	fmt.Println("--------------添加客户--------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name,gender,age,phone,email)
	if cv.customerController.Add(customer) {
		fmt.Println("--------------添加成功--------------\n")
	} else{
		fmt.Println("--------------添加失败--------------\n")
	}
}

//得到输入，删除该id对应的客户
func (cv *customerView) delete() {
	fmt.Println("--------------删除客户--------------")
	fmt.Print("客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	if cv.customerController.Delete(id) {
		fmt.Println("--------------删除成功--------------\n")
	} else {
		fmt.Println("--------------查无此人--------------\n")

	}
}

func (cv *customerView) update() {
	fmt.Println("--------------修改信息--------------")
	fmt.Print("客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := cv.customerController.FindById(id)
	if index == -1 {
		fmt.Println("--------------查无此人--------------\n")
		return
	}
	fmt.Printf("姓名(%v):",cv.customerController.Customers[index].Name)
	name := ""
	fmt.Scanln(&name)

	fmt.Printf("性别(%v):",cv.customerController.Customers[index].Gender)
	gender := ""
	fmt.Scanln(&gender)

	fmt.Printf("年龄(%v):",cv.customerController.Customers[index].Age)
	age := 0
	fmt.Scanln(&age)

	fmt.Printf("电话(%v):",cv.customerController.Customers[index].Phone)
	phone := ""
	fmt.Scanln(&phone)

	fmt.Printf("邮箱(%v):",cv.customerController.Customers[index].Email)
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(id,name,gender,age,phone,email)

	cv.customerController.Update(customer)
	fmt.Println("--------------修改成功--------------\n")
}