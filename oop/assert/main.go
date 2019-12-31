package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point=Point{1,2}
	a = point
	//将a赋值给一个point变量
	var b Point
	//b =a 不可以
	//b=a.(Point) 可以
	b=a.(Point)
	fmt.Println(b)

	//类型断言其他案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2
	y:=x.(float32)
	fmt.Printf("y type is %T, value is %v",y,y)
}
