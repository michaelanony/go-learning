package main

import (
	"../service"
	"fmt"
)

type customerView struct{
	key string
	loop bool
	customerService *service.CustomerService
}
func(this *customerView) list(){
	customers := this.customerService.List()
	for i :=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
}
func (this *customerView) mainMenu(){
	for {
		fmt.Println("------Manager--------")
		fmt.Println("1.Add Customer")
		fmt.Println("2.Change Customer")
		fmt.Println("3.Delete Customer")
		fmt.Println("4.Customer Lists")
		fmt.Println("5.Quit")
		fmt.Println("Please choose option")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println()
		case "2":
			fmt.Println()
		case "3":
			fmt.Println()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("Wrong type")
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("Quit success!")
}
func main() {
	customerView := customerView{
		key:"",
		loop:true,
	}
	customerView.customerService = service.NewCustomerSvc()
	customerView.mainMenu()

}
