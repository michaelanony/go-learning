package controller

import "github.com/gin-gonic/gin"

func IndexHandle(c *gin.Context)  {

}
func RegistryHandle(c *gin.Context)  {

}
func LoginHandle(c *gin.Context)  {
	c.DefaultPostForm("type")
}