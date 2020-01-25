package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	//1、重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})
	//2、异步
	r.GET("/long_async", func(c *gin.Context) {
		//新建副本
		copyContent := c.Copy()
		//异步处理
		go func(){
			time.Sleep(3 * time.Second)
			log.Println("异步执行："+copyContent.Request.URL.Path)
		}()
	})
	//3、同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行："+c.Request.URL.Path)
	})
	r.Run(":8000")
}
