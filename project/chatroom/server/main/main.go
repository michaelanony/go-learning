package main

import (
	"fmt"
	"go-learning/project/chatroom/server/model"
	"io"
	"net"
)



//编写ServerProcessMes函数
//功能：根据客户端发送的消息种类不同，决定调用哪个函数来处理

func process(conn net.Conn)  {
	defer conn.Close()
	processor := &Processor{
		Conn:conn,
	}
	err:=processor.process()
	if err!=nil{
		fmt.Println(err)
	}
}
//完成对UserDao的初始化任务
func initUserDao()  {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	initPool("192.168.11.31:30002")
	initUserDao()//需要在initPool后，因为需要pool
	fmt.Println("Server is listening on 8889...")
	listen,err:=net.Listen("tcp","127.0.0.1:8880")
	defer listen.Close()
	if err!=nil{
		panic(err)
		return
	}
	for {
		fmt.Println("Waiting for client to connect...")
		conn,err := listen.Accept()
		if err!=nil{
			if err ==io.EOF{
				fmt.Println("对方关闭了链接,服务端退出")
				return
			}
		}
		//一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
