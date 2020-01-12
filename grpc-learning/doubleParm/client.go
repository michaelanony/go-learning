package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err.Error())
	}

	var resp *float32 //返回值

	param := &AddParma{Args1: 1.2, Args2: 2.3}
	//同步调用
	err = client.Call("MathUtil.CalTwo", param, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)

}
