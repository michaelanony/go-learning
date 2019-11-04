package main

import "fmt"

var alex= "dsb"

func main(){
	var name string;
	var age int;
	fmt.Println(name)
	fmt.Println(age)
	var(
		a string
		b int
		c bool
	)
	a="a"
	b=100
	c=true
	fmt.Println(a,b,c)

	var x string = "boy"
	fmt.Printf("%s,HaHa,%d",x,b)
	// 类型推导
	var y =200
	var z = true
	fmt.Println(y,z)

	//简短变量声明（只能在函数内部使用）
	me := "my"
	fmt.Println(me)
}