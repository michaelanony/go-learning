package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context)  {
	
}
func main() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		//url参数可以通过DefaultQuery（）或Query（）方法获取
		//DefaultQuery（）若参数不存在，返回默认值，Query（）若不存在，返回空串
		name :=c.DefaultQuery("name","jack")//第二个参数是默认值
		//name := c.Param("name")
		action:=c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.POST("/posttest",getting)
	r.PUT("/puttest")
	r.Run(":8000")
}
