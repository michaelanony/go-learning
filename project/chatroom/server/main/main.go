package main

import (
	"fmt"
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
		panic(err)
	}
}
func main() {
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
