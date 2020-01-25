package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//定义程序计时中间件，然后定义2个路由，执行函数后应该打印统计的执行时间
func myTime(c *gin.Context)  {
	start:=time.Now()
	c.Next()
	since:=time.Since(start)
	fmt.Println("Spend time:",since)
}

func main() {
	r :=gin.Default()
	r.Use(myTime)
	sg := r.Group("/shopping")
	{
		sg.GET("/index",shopIndexHandler)
	}
	r.Run(":8000")
}

func shopIndexHandler(c *gin.Context)  {
	time.Sleep(2*time.Second)

}