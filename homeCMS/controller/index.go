package controller

import (
	"github.com/gin-gonic/gin"
	"go-learning/homeCMS/controller/fileHandle"
	"go-learning/homeCMS/controller/userHandle"
)

func GinRouter(r *gin.Engine) *gin.Engine {
	rr :=r.Group("/")
	userHandle.Routers(rr)
	fileHandle.Routers(rr)
	return r
}
