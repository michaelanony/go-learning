package process

import (
	"fmt"
	"go-learning/project/chatroom/client/utils"
	"net"
	"os"
)

func ShowMenu()  {
	fmt.Println("----Login Success----")
	fmt.Println("----1、显示在线用户列表----")
	fmt.Println("----2、发送消息----")
	fmt.Println("----3、信息列表----")
	fmt.Println("----4、退出系统----")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("正在退出系统...")
		os.Exit(0)
	}
}

func serverProcessMes(conn net.Conn)  {
	tf := &utils.Transfer{
		Conn:conn,
	}
	for{
		fmt.Println("Waiting for message from server")
		mes,err:=tf.ReadPkg()
		if err !=nil{
			fmt.Println("Read message error")
			panic(err)
		}
		fmt.Printf("mes=%v",mes)
	}
}