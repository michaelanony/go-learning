package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func main() {

	testUrl := ""
	start := time.Now()
	defer func(time1 time.Time) {
		log.Printf("cost Seconds:%v\n", time.Since(time1))
	}(start)
	//handleUrl(testUrl)
	regexUrl(testUrl)
}

//500-800
func handleUrl(testUrl string) {
	for i := 0; i <= 10000; i++ {
		ret, _ := url.Parse(testUrl)
		ret1 := strings.Split(ret.Path, "/")
		ret2 := strings.Split(ret1[len(ret1)-1], "-")
		values, _ := url.ParseQuery(ret.RawQuery)
		fmt.Println(values["trid"][0])
		fmt.Println(values["platform"][0])
		fmt.Println(ret2[2])
	}
}

//250-604
func regexUrl(testUrl string) {
	ptrid, _ := regexp.Compile("trid=([0-9a-zA-Z]+)&")
	pplat, _ := regexp.Compile("platform=(.*?)&")
	pqn, _ := regexp.Compile("(\\-[0-9]+\\-)([0-9]+\\.(flv|m4s|mp4))")
	for i := 0; i <= 10000; i++ {
		find := ptrid.FindString(testUrl)
		find1 := pplat.FindString(testUrl)
		find2 := pqn.FindString(testUrl)
		fmt.Println(find)
		fmt.Println(find1)
		fmt.Println(find2)
	}
}
