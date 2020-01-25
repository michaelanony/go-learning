package main
//模拟实现权限验证中间件
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并检验
		if cookie,err:=c.Cookie("abc");err==nil{
			if cookie =="123"{
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized"})
		//若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	r :=gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//设置cookie
		c.SetCookie("abc","123",60,"/","127.0.0.1",false,false)
		c.String(200,"Login Success")
	})
	r.GET("/home",AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200,gin.H{"data":"home"})
	})
	r.Run(":8000")
}
