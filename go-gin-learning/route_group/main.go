package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	v1:=r.Group("/v1")
	//大括号是书写规范
	{
		v1.GET("/login",login)
		v1.GET("/submit",submit)
	}
	v2:=r.Group("/v2")
	{
		v2.POST("/login",login)
		v2.POST("/submit",submit)
	}
	if err:=r.Run(":8000");err!=nil{
		panic(err)
	}

}
func login(c *gin.Context) {
	name :=c.DefaultQuery("name","jack")//第二个参数是默认值
	c.String(http.StatusOK, "Hello %s",name)
}
func submit(c *gin.Context)  {
	return
}