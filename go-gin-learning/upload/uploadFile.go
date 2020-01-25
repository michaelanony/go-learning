package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制表单上传大小 8mb，默认为32mb
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload",func(c *gin.Context){
		form ,err:=c.MultipartForm();
		if err!=nil{
			c.String(http.StatusBadRequest,fmt.Sprintf("get err %s"),err.Error())
		}
		files := form.File["files"]
		for _,file := range files{
			if err:=c.SaveUploadedFile(file,file.Filename);err!=nil{
				c.String(http.StatusBadRequest,fmt.Sprintf("upload err %s"),err.Error())
				return
			}
		}
		////从表单取文件
		//file,_:=c.FormFile("file")
		//log.Println(file.Filename)
		////传到项目根目录，名字就用本身的
		//c.SaveUploadedFile(file,file.Filename)
		c.String(200,fmt.Sprintf("You have %d upload",len(files)))
	})
	r.Run(":8000")
}
