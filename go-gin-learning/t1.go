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
		c.DefaultQuery()
		name := c.Param("name")
		action:=c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.POST("/posttest",getting)
	r.PUT("/puttest")
	r.Run(":8000")
}
