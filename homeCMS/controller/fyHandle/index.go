package fyHandle

import (
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/service"
)

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/fy")
	rr.GET("/all",service.GetStaticNumSvc)
	rr.GET("/current",service.GetCurrentCityStatus)
	rr.GET("/citylist",service.GetCityList)
}


