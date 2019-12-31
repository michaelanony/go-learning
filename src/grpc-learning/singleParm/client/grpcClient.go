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
	var req float32 //请求值
	req = 3
	var resp *float32 //返回值

	//同步调用
	//err = client.Call("MathUtil.Cal",req,&resp)
	//if err!=nil{
	//	panic(err.Error())
	//}
	//fmt.Println(*resp)

	//异步调用
	syncCall := client.Go("MathUtil.Cal", req, &resp, nil)
	replayDone := <-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(*resp)
}
