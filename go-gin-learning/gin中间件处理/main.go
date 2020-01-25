package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//gin可以构建中间件，但他只对注册过的路由函数起作用
//对于分组路由，嵌套使用中间件，可以限定中间件的作用范围
//中间件分为全局中间件，单个路由中间件和群组中间件
//gin中间件必须是一个gin.HandlerFunc类型

//定义中间件
func MiddleWare()gin.HandlerFunc  {
	return func(c *gin.Context) {
		t:=time.Now()
		fmt.Println("中间件开始执行了")
		//设置变量到context的key中，可以通过Get（）获取
		c.Set("request","middle ware")
		//执行中间件
		c.Next()
		//中间件执行完后续的一些事情
		status :=c.Writer.Status()
		fmt.Println("中间件执行完毕",status)
		t2:=time.Since(t)
		fmt.Println("time:",t2)
	}
}
func main() {
	r := gin.Default()
	//注册中间件
	r.Use(MiddleWare())//申明全局中间件
	//大括号是为了代码规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			//取值
			req,_:=c.Get("request")
			fmt.Println("request:",req)
			c.JSON(200,gin.H{"request":req})
		})
		//跟在路由后面，使用局部中间件
		r.GET("/partMiddleware", MiddleWare(),func(c *gin.Context) {
			//取值
			req,_:=c.Get("request")
			fmt.Println("request:",req)
			c.JSON(200,gin.H{"request":req})
		})
	}
	r.Run(":8000")
}
