package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()

	//1、返回json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"success","status":"OK"})
	})
	//2、结构体响应
	r.GET("/struct", func(c *gin.Context) {
		var msg struct{
			Name string
			Age int
			Message string
		}
		msg.Name="root"
		msg.Message="msg"
		msg.Age=23
		c.JSON(200,msg)
	})

	//3、xml
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200,gin.H{"message":"test"})
	})

	//4、yaml
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200,gin.H{"message":"yaml"})
	})

	//5、protoBuf高效存储读取的工具
	r.GET("/protoBuf", func(c *gin.Context) {
		resp:=[]int64{int64(1),int64(2)}
		label:="label"
		data:=&protoexample.Test{
			Label:&label,
			Reps:resp,
		}
		c.ProtoBuf(200,data)
	})
	r.Run(":8000")
}
