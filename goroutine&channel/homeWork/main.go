package main

import "fmt"

//统计1-200000的数字中，哪些是素数？

func putNum(intChan chan int){
	for i:=1;i<=10000;i++{
		intChan<-i
		fmt.Printf("Now put num:%d",i)
	}
	close(intChan)
}
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool){
	var flag bool
	for{
		num,ok := <-intChan
		if !ok{
			break
		}
		flag = true
		for i:=2;i<num;i++{
			if num % i == 0{
				flag = false
				break
			}
		}
		if flag{
			primeChan<-num
		}
	}
	exitChan<-true
}
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)
	go putNum(intChan)
	for i:=0;i<4;i++{
		go primeNum(intChan,primeChan,exitChan)
	}
	go func(){
	for i :=0;i<4;i++{
		<-exitChan
	}
	close(primeChan)
	}()
	for{
		res,ok :=<-primeChan
		if !ok{
			break
		}
		fmt.Printf("The num = %d\n",res)
	}

}