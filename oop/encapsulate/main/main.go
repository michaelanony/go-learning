package main
import (
	"../model"
	"fmt"
)
func main() {
	p:=model.NewPerson("smith")
	p.SetAge(18)
	p.SetSalary(2000)
	fmt.Println(p)
}