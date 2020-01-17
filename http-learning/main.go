package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Test struct {
}

func get(wrongChan chan string, exitChan chan bool) {
	for i := 0; i < 100; i++ {
		url := ""
		cookie:=&http.Cookie{}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		ret, _ := ioutil.ReadAll(resp.Body)
		strRet := string(ret)

		if strings.Contains(strRet, "\"data\":null") {
			fmt.Println(strRet)
			wrongChan <- strRet
		}
	}
	exitChan <- true

}

func main() {
	wrongChan := make(chan string, 1000)
	exitChan := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go get(wrongChan, exitChan)
	}
	go func() {
		for i := 0; i < 100; i++ {
			<-exitChan
		}
		close(wrongChan)
	}()
	for {
		res, ok := <-wrongChan
		if !ok {
			break
		}
		fmt.Println(res)
	}

}
