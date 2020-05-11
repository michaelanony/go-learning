package main

//1、需要监听
//2、需要实例化gRPC服务端
//3、在gRPC上注册为服务
//4、启动服务端

import (
	"context"
	"fmt"
	pb "go-learning/go-micro-learning/grpc-demo/proto"
	"google.golang.org/grpc"
	"net"
)

//定义空接口
type UserInfoService struct {}

var u = UserInfoService{}

//实现方法
func (s *UserInfoService)GetUserInfo(ctx context.Context,req *pb.UserRequest) (resp *pb.UserResponse,err error) {
	//通过用户名查询用户信息
	name := req.Name
	if name =="zs"{
		resp = &pb.UserResponse{
			Id: 1,
			Name: name,
			Age: 22,
			Hobby: []string{"sing","Run"},
		}
	}
	return
}

func main() {
	//1、监听
	//地址
	addr:="127.0.0.1:8080"
	listener,err:=net.Listen("tcp",addr)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Listening...")
	//2、实例化gRPC
	s:=grpc.NewServer()
	//3、在gRPC上注册微服务
	pb.RegisterUserInfoServiceServer(s,&u)
	//4、启动服务端
	if err=s.Serve(listener);err!=nil{
		panic(err)
	}
}