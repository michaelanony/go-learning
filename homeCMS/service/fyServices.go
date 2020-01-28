package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/dao"
	"go-learning/homeCMS/model"
)

func GetStaticNumSvc(c *gin.Context){
	sn := &model.StaticNum{}
	sn, err := dao.GinDao.GetFyStaticLatest()
	if err!=nil{
		fmt.Println(err)
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":sn,
	})
}
func GetCurrentCityStatus(c *gin.Context)  {
	param :=c.DefaultQuery("param","all")
	city :=c.DefaultQuery("city","all")
	cs, err := dao.GinDao.GetCurrentCityStatus(city, param)
	if err!=nil{
		fmt.Println(err)
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"data":cs,
	})
}