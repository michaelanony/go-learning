package service
import ("../model")
//增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段。表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

func NewCustomerSvc() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum=1
	customer:= model.NewCustomer(1,"michael","male",20,"112","1@q.com")
	customerService.customers = append(customerService.customers,customer)
	return customerService
}

func (this *CustomerService)  List() []model.Customer{
	return this.customers
}