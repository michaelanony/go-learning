package main

import (
	"fmt"
	"strconv"
	"time"
)

func goroutine() {
	for i := 1; i <= 10; i++ {
		fmt.Println("goroutine() hello" + strconv.Itoa(i))
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
	//num := runtime.NumCPU()
	//runtime.GOMAXPROCS(num)
	//fmt.Println("num=", num)
	a := make(chan int, 3)
	a <- 1
	a <- -1
	a <- 2
	close(a)
	for v := range a {
		fmt.Println(v)
	}

}
