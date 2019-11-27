package main

import (
	"fmt"
)

type Account struct {
	Account string
	Pwd string
	Balance float64
}

func (account *Account) Deposite(money float64,pwd string)  {
	if pwd !=account.Pwd{
		fmt.Println("Wrong pwd")
		return
	}
	if money<=0{
		fmt.Println("Wring money")
		return
	}
	account.Balance+=money
	fmt.Println("Succeed!")
}
func (account *Account) WithDraw(money float64,pwd string)  {
	if pwd !=account.Pwd{
		fmt.Println("Wrong pwd")
		return
	}
	if money<=0|| money>account.Balance{
		fmt.Println("Wring money")
		return
	}
	account.Balance -= money
	fmt.Println("Take money succeed!")
}
func (account *Account) Query(pwd string)  {
	if pwd !=account.Pwd{
		fmt.Println("Wrong pwd")
		return
	}
	fmt.Printf("Your account is:%v,money=%v\n",account.Account,account.Balance)
}
func main() {
	account := Account{
		Account: "gs",
		Pwd:     "123",
		Balance: 666,
	}
	account.Query("123")
	account.Deposite(124141,"123")
	account.Query("123")
	account.WithDraw(222,"123")
	account.Query("123")
}
