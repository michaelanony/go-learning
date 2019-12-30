package main

import (
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct{

}
func (mu *MathUtil)Cal(req float32,resp *float32) error{
	*resp = req+req
	return nil
}
	func main() {
	//1、初始化指针类型
	ma := new(MathUtil)
	//2、注册服务对象
	err:=rpc.Register(ma)
	if err!=nil{
		panic(err.Error())
	}
	//3、将mathutil提供的服务注册到http上
	rpc.HandleHTTP()
	//4、在特定端口监听
	listen,err:=net.Listen("tcp","127.0.0.1:8081")
	if err!=nil{
		panic(err.Error())
	}
	http.Serve(listen,nil)

	}
