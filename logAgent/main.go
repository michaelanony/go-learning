package main

import (
	"fmt"
	"go-learning/logAgent/kafka"
	"go-learning/logAgent/taillog"
	"time"
)

//logAgent 入口程序

func run()  {
	for {
		select{
		case line:=<- taillog.ReadLog():
			kafka.SendToKafka("web_log",line.Text)
		default:
			time.Sleep(time.Second)
		}

	}
}
func main() {
	//1、加载配置文件


	//2.初始化kafka连接
	err:=kafka.Init([]string{"192.168.11.3:9092"})
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Kafka init success!")
	//3、打开日志文件准备收集日志
	err = taillog.Init("./t.log")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Tail init success!")
	run()
}