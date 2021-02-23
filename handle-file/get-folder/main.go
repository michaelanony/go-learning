package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const TargetPath = "/Users/michael/tmp/newDomain1"

type Cal5MinTarget struct {
	FilePath string
	Domain   string
	Country  string
}
type Cal5MinArg []Cal5MinTarget
type Cal5MinArgs []Cal5MinArg

func main() {
	rootDir, err := ioutil.ReadDir(TargetPath)
	if err != nil {
		log.Println("read root dir failed")
		return
	}
	var cmas Cal5MinArgs
	for _, domainPath := range rootDir {
		var cma Cal5MinArg
		if domainPath.IsDir() {
			domain := domainPath.Name()
			fmt.Println(domain)
			datePaths, err := ioutil.ReadDir(TargetPath + "/" + domainPath.Name())
			if err != nil {
				log.Println("read datePath dir failed")
				return
			}
			for _, datePath := range datePaths {
				if !datePath.IsDir() && strings.HasSuffix(datePath.Name(), ".gz") {
					fmt.Println(datePath.Name())
					tmp := &Cal5MinTarget{
						FilePath: datePath.Name(),
						Domain:   domain,
						Country:  "all",
					}
					cma = append(cma, *tmp)
				}
			}
		}
		if len(cma) > 0 {
			cmas = append(cmas, cma)
		}
	}
	return
}
