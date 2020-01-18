package main

import (
	"fmt"
	"go-learning/project/chatroom/common/message"
	processes "go-learning/project/chatroom/server/process"
	"go-learning/project/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登入逻辑
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理！")
	}
	return
}

func (this *Processor) process() (err error) {
	for {
		//这里我们将读取数据包，直接封装成一个函数，返回message
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client quit.")
				return err
			} else {
				panic(err)
				return err
			}
		}
		fmt.Println("message=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			panic(err)
		}
	}
}
