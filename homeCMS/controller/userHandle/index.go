package userHandle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/errno"
	"go-learning/homeCMS/model"
	"go-learning/homeCMS/service"
	"net/http"
)

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/user")
	rr.GET("/login", LoginHandle)
	rr.POST("/registry", RegistryHandle)
}
func RegistryHandle(c *gin.Context)  {
	ip := c.ClientIP()
	fmt.Println(ip)
	user :=&model.HomeUser{}
	user.UNickname=c.PostForm("u_nickname")
	user.UName=c.PostForm("u_name")
	user.UPassword=c.PostForm("u_password")
	user.URegisterIp=c.ClientIP()
	err := service.RegistrySvc(user)
	if err!=nil{
		if err==errno.ERROR_USER_EXISTS{
			c.JSON(http.StatusOK,gin.H{
				"code":300,
				"msg":"用户已经存在",
			})
			return
		}else if err ==errno.ERROR_USER_FORMAT{
			c.JSON(http.StatusOK,gin.H{
				"code":300,
				"msg":"格式不正确！",
			})
			return
		}

	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"Register success!",
	})
}
func LoginHandle(c *gin.Context)  {

}