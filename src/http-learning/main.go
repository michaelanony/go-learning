package main

import (
"fmt"
"net/http"
)

const domainUrl  = "http://callback-api.bilibili.cn"
func get(url string){
	client := &http.Client{}
	targetUrl := domainUrl+url
	req,_:=http.NewRequest("GET",targetUrl,nil)
	rep,_ := client.Do(req)

}
func main() {
	get("/x/tv/modpage_v2")
}
