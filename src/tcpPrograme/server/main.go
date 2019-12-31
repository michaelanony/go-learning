package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start listening...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	fmt.Printf("listen suc=%v\n", listen)
}
