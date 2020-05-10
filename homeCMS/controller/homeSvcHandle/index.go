package homeSvcHandle
import (
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/service"
)

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/home")
	rr.GET("/hostsconfig",service.GetHomeHosts)

}

