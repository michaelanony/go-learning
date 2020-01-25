package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/controller"
	"go-learning/homeCMS/dao/db"
)

func main() {
	router :=gin.Default()
	rr := controller.GinRouter(router)
	mysqlDns:="michael:cccbbb@tcp(192.168.11.31:30001)/testDb?parseTime=true"
	err := db.Init(mysqlDns, "")
	if err!=nil{
		panic(err)
	}
	_, err = db.GetUser("michael")
	if err!=nil{
		fmt.Println(err)
	}

	if err=rr.Run("0.0.0.0:80");err!=nil{
		panic(err)
	}
}
