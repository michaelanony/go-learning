package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/controller"
	"go-learning/homeCMS/dao"
	"go-learning/homeCMS/middleWare"
)


func main() {
	if err:= dao.InitPool("michael:cccbbb@tcp(192.168.11.31:30001)/testDb?parseTime=true","192.168.11.31:30002");err!=nil{
		panic(err)
	}
	router :=gin.Default()
	router.Use(middleWare.EnableRedisSession(),middleWare.Cors())
	rr := controller.GinRouter(router)
	if err:=rr.Run(":80");err!=nil{
		panic(err)
	}
}
