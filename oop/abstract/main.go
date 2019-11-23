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
}
func main() {

}
