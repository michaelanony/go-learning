package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func goroutine(){
	for i:=1;i<=10;i++{
		fmt.Println("goroutine() hello"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	//go goroutine()//开启一个协程
	//for i:=1;i<=10;i++{
	//	fmt.Println("main() hello"+ strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}
	//获取当前cpu数量
	num:= runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println("num=",num)
}