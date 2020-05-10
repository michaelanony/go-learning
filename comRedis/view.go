package main

import (
	"fmt"
	"os"
)


func view()  {
	var key int
	for {
		fmt.Println("—————————请选择获取验证码方法——————————")
		fmt.Println("\t 1 通过captcha_key")
		fmt.Println("\t 2 通过手机号码")
		fmt.Println("\t 3 退出系统")
		fmt.Println("\t 请选择（1-3）")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("请输入captcha_key:")
			fmt.Scanf("%s\n",&captchaKey)
			if err:=UserDao.GetCodeByCaptchaKey(captchaKey);err!=nil{
				fmt.Println(err)
			}


		case 2:
			fmt.Println("请输入手机号码:")
			fmt.Scanf("%s\n",&phoneNum)
			fmt.Println(phoneNum)

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误")
		}

	}
}