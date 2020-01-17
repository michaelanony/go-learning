package main

import (
	"fmt"
	"go-learning/project/chatroom/common/message"
	"net"
)

func serverProcessMes(conn net.Conn,mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登入逻辑
		err = serverProcessLogin(conn,mes)
	case message.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理！")
	}
	return
}