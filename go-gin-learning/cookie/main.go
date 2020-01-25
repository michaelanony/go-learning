package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r :=gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		//获取客户端是否携带cookie
		cookie,err:=c.Cookie("key_cookie")
		//err不为空时说明cookie不存在
		if err!=nil{
			fmt.Println(cookie)
			cookie = "NotSet"
			//设置cookie
			//第三个选项为过期时间，单位为秒,第四个为所在目录，第五个为domain
			//第六个为是否只能通过http是访问，第七个是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie","value_cookie",60,"/","127.0.0.1",false,false)
		}
		fmt.Println("cookie值为",cookie)
	})
	r.Run(":8000")
}
