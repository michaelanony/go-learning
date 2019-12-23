package main

import "fmt"

func writeData(intChan chan int){
	for i:=0;i<50;i++{
		intChan<-i
	}
	close(intChan)
}

func readData(intChan chan int,exitChan chan bool){
	for {
		v,ok:=<-intChan
		if !ok{
			break
		}
		fmt.Printf("readData =%v\n",v)
	}
	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}
func main() {
	//创建两个管道
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)
	go writeData(intChan)
	go readData(intChan,exitChan)
}