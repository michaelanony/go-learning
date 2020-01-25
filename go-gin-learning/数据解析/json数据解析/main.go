package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`//binding表示为必须字段，使用required时若接受为空值则报错
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}
func main() {
	r := gin.Default()
	var jsonText User

	//1、绑定json数据
	r.POST("/bindJson", func(ctx *gin.Context) {
		if err:=ctx.ShouldBindJSON(&jsonText);err!=nil{
			//返回错误信息
			//gin.H封装了生成json数据的工具
			ctx.JSON(http.StatusBadRequest,gin.H{"ERROR":err.Error()})
			return
		}
		if jsonText.Username !="root" || jsonText.Password!="root"{
			ctx.JSON(http.StatusBadRequest,gin.H{"status":403})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"status":200})
	})

	//2、绑定表单数据
	r.POST("/bindForm", func(ctx *gin.Context) {
		if err:=ctx.Bind(&jsonText);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if jsonText.Username !="root" || jsonText.Password!="root"{
			ctx.JSON(http.StatusBadRequest,gin.H{"status":403})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"status":200})
	})

	//3、绑定uri数据
	r.GET("/bindUri/:username/:password", func(ctx *gin.Context) {
		if err:=ctx.ShouldBindUri(&jsonText);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if jsonText.Username !="root" || jsonText.Password!="root"{
			ctx.JSON(http.StatusBadRequest,gin.H{"status":403})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"status":200})
	})
	if err:=r.Run(":8000");err!=nil{
		panic(err)
	}
}
