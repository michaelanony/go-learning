package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const fn = "/Users/michael/Desktop/upos-sz-mirrorkodob.bilivideo.com_2021-01-09-02_part-00001.gz"

type target struct {
}

func main() {
	cal("tx")
}
func cal(cdn string) {
	fr, err := os.Open(fn)
	if err != nil {
		return
	}
	defer fr.Close()
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return
	}
	defer gr.Close()
	var (
		//aborLen     = 0
		bucketError = 0
		sumCnt      = 0
		hitCnt      = 0
		//hitError    = 0
	)
	rows := bufio.NewScanner(gr)
	for rows.Scan() {
		line := rows.Text()
		split := strings.Split(line, "|")
		rawUrl := split[4]
		fmt.Println(rawUrl)
		netProto := strings.Split(rawUrl, "/")[0]
		bucket := ""
		if strings.Contains(strings.ToLower(cdn), "xunlei") {
			bucket = strings.Split(rawUrl, "/")[4]
		} else if strings.Contains(netProto, "http") || strings.Contains(netProto, "https") {
			bucket = strings.Split(rawUrl, "/")[1]
		} else {
			bucketError++
		}
		parseUrl, err := url.Parse(rawUrl)
		if err != nil {
			return
		}
		urlPath := strings.Split(parseUrl.Path, "/")
		urlPathParse := strings.Split(urlPath[len(urlPath)-1], "-")
		querys, _ := url.ParseQuery(parseUrl.RawQuery)
		trid := querys["trid"][0]
		platform := querys["platform"][0]
		qn := urlPathParse[2]
		//TODO
		//cfg map
		if bucket == "upgcxcode" && len(trid) == 33 {
			bucket = bucket + "_" + string(trid[32])
		}
		urlStr := platform + "#" + bucket + qn
		hit := split[8]
		if strings.ToUpper(hit) == "HIT" {
			hitCnt = 1
		}
		sumCnt = 1
		key := split[2]
		bytes := split[9]
		key = key[0:15] + strconv.ParseInt(key[15:16]/5*5)

	}
}
