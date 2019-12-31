package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.De
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	r.Run()
}
