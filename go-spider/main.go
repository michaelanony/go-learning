package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HttpGet(url string) (res string ,err error) {
	resp, err := http.Get(url)
	if err!=nil{
		return
	}
	defer resp.Body.Close()
	buf:=make([]byte) 4096
	for {
		n, err := resp.Body.Read(buf)
		if n==0{
			fmt.Printf()
		}
	}
}
func working(start, end int) {
	fmt.Printf("从%d到%d", start, end)
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E8%B4%B4%E5%90%A7&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		ret, err := HttpGet(url)
		if err != nil {
			log.Println("http err:", err)
			continue
		}

	}
}
func main() {
	var start, end int
	fmt.Println("请输入爬取的起始页")
	fmt.Scan(&start)
	fmt.Println("end:")
	fmt.Scan(&end)

}
