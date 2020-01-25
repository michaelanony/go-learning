package main

import (
"fmt"
"github.com/go-log/log"
"github.com/micro/go-micro"
"go-learning/go-micro-learning/example1/proto/sum"
"go-learning/go-micro-learning/example1/sum-srv/handler"
)

func main() {
	//创建服务
	srv:=micro.NewService(
		micro.Name("go.micro.learning.srv.sum"), )
	//初始化
	srv.Init(
		micro.BeforeStart(
			func() error {
				fmt.Println("sss")
				log.Log("[srv]启动前日志")
				return nil
			}),
		micro.AfterStart(
			func() error {
				log.Log("[srv]启动后日志")
				return nil
			}),
	)
	//挂载接口
	_ = sum.RegisterSumHandler(srv.Server(),handler.Handler())
	if err:=srv.Run();err!=nil{
		panic(err)
	}
}
