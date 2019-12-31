package main

import "fmt"

const pi = 3.14

//批量声明常量
const (
	a = 100
	b

	c = 1000
	d
	e
)

//iota
const (
	aa = iota //0
	bb        //1
	_         //2
	dd        //3
)
const (
	n1 = iota //0
	n2 = 100  //100
	n3 = iota //2
	n4        //3
)
const n5 = iota //0

const (
	_  = iota
	KB = 1 << (10 * iota) //s
	MB = 1 << (10 * iota)
)

const (
	a1, b1 = iota + 1, iota + 2 //iota=0;1,2
	c1, d1                      //iota=1;2,3
	e1, f1                      //iota=2;3,4
)

func main() {

	fmt.Println(pi, a, b, c, d, e)
	fmt.Println(aa, bb, dd)

}

//const多声明中，若不写，就默认和上行一样
//iota
//1.遇到const iota初始化为0
//2.const中每新增一行，iota就递增1
