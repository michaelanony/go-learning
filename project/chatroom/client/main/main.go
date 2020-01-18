package main

import (
	"fmt"
	"go-learning/project/chatroom/client/process"
	"os"
)

var userId int
var userPwd string


func main() {
	var key int
	for {
		fmt.Println("——————————欢迎登入多人聊天系统——————————")
		fmt.Println("\t\t\t 1 登入聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")

		fmt.Scanf("%d\n", &key)
		switch key {
			case 1:
				fmt.Println("登入聊天室")
				fmt.Println("请输入ID号：")
				fmt.Scanf("%d\n", &userId)
				fmt.Println("请输入密码：")
				fmt.Scanf("%s\n", &userPwd)
				up := &process.UserProcess{}
				up.Login(userId,userPwd)

			case 2:
				fmt.Println("注册用户")
			case 3:
				fmt.Println("退出系统")
				os.Exit(0)
			default:
				fmt.Println("你的输入有误")
		}

	}
}
