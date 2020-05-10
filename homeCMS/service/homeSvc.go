package service

import (
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/dao"
	"log"
)

func GetHomeHosts(c *gin.Context)  {
	hf ,err :=dao.GinDao.GetHostsConfig()
	if err !=nil{
		log.Println(err)
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":hf,
	})
}