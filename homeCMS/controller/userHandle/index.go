package userHandle

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/dao"
	"go-learning/homeCMS/errno"
	"go-learning/homeCMS/model"
	"go-learning/homeCMS/service"
	"net/http"
)

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/api/user")
	rr.POST("/login", LoginHandle)
	rr.POST("/registry", RegistryHandle)
	rr.GET("/test",CurrentUserTest)
}
func RegistryHandle(c *gin.Context)  {
	ip := c.ClientIP()
	if exist,_ :=dao.GinDao.CheckIpInRedis(ip);exist{
		c.JSON(http.StatusOK,gin.H{
			"code":300,
			"msg":"请不要频繁注册",
		})
		return
	}
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
	_ = dao.GinDao.InsetIntoRedis(ip)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"Register success!",
	})
}
func LoginHandle(c *gin.Context)  {
	user :=&model.HomeUser{}
	user.UNickname=c.PostForm("u_nickname")
	user.UPassword=c.PostForm("u_password")
	user ,err:= dao.GinDao.GetUser(user.UNickname,user.UPassword)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":401,
			"msg":"账号密码错误！",
		})
		return
	}
	session:=sessions.Default(c)
	session.Set("username",user.UNickname)
	session.Save()
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"登入成功！",
		"userInfo":user,
	})
}

func CurrentUserTest(c *gin.Context)  {
	session := sessions.Default(c)
	sessValue := session.Get("username")
	username,ok := sessValue.(string)
	if !ok{
		return
	}
	c.String(200,username)
}