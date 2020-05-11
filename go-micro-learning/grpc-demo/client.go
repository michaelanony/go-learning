package main

//1、连接服务端
//2、实例化gRPC客户端
//3、调用

import (
	"context"
	"fmt"
	pb "go-learning/go-micro-learning/grpc-demo/proto"
	"google.golang.org/grpc"
)

func main() {
	//1、连接服务端
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//2、实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	//3、组装请求参数
	req := new(pb.UserRequest)
	req.Name="zs"
	//3、调用
	response, err := client.GetUserInfo(context.Background(), req)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("resp:%v\n",response)
}
