package main

import "fmt"

var alex= "dsb"

//定义有两个返回值的函数
func foo()(string,int){
	return "alex",1000
}

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

	//调用foo函数
	// _ 用于接受不需要的变量值
	aa, _ := foo()
	_,bb:=foo()
	fmt.Println(aa,bb)
	
	//不能重复声明同名变量
	//var aname = "Tom"
	var aname = "Michael"
	fmt.Println(aname)

}