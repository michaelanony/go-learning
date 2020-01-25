package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/form",func(c *gin.Context){
		c.DefaultPostForm("type","alert")
		username:=c.PostForm("userHandle")
		password := c.PostForm("password")
		c.PostFormArray("")//勾选框
		c.String(http.StatusOK, username+" is "+password)
	})
	r.Run(":8000")
}
