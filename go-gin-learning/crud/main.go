package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	err := InitDb()
	if err!=nil{
		panic(err)
	}
	r :=gin.Default()
	bp := r.Group("/book")
	{
		bp.GET("/list",bookListHandler)
	}
	r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList,err:=QueryAllBook()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":300,
			"msg":err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"data":bookList,
	})
}