package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/controller"
	"go-learning/homeCMS/dao/db"
)

func main() {
	router :=gin.Default()
	mysqlDns:="michael:cccbbb@tcp(192.168.11.31:30001)/testDb?parseTime=true"
	err := db.Init(mysqlDns, "")
	if err!=nil{
		panic(err)
	}
	_, err = db.GetUser("michael")
	if err!=nil{
		fmt.Println(err)
	}
	router.GET("/login",controller.LoginHandle)
	router.POST("/registry",controller.RegistryHandle)
	if err=router.Run(":8000");err!=nil{
		panic(err)
	}
}
