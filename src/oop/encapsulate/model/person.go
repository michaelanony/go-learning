package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

//factory function
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和salary 编写get，set方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("Wrong age ")
	}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetSalary(sal float64) {
	if sal < 0 {
		fmt.Println("Wrong salary")
	} else {
		p.salary = sal
	}
}
func (p *person) GetSalary() float64 {
	return p.salary
}
