package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"go-learning/go-micro-learning/example1/proto/sum"
	"strconv"
)

var (
	srvClient sum.SumService
)

type Open struct {

}

func (open Open) Fetch(ctx context.Context,req *api.Request,rsp *api.Response) error  {
	sumInputStr:=req.Get["sum"].Values[0]
	sumInput,_:=strconv.ParseInt(sumInputStr,10,10)
	sumReq:=&sum.SumRequest{
		Input:sumInput,
	}
	sumRsp ,_:=srvClient.GetSum(ctx,sumReq)
	ret,_:=json.Marshal(map[string]interface{}{
		"sum":sumRsp,
	})
	rsp.Body =string(ret)
	return nil
}
func main() {
	service :=micro.NewService(
		micro.Name("go.micro.learning.api.open"),)
	srvClient = sum.NewSumService("go.micro.learning.srv.sum",service.Client())
	//暴露接口
	//todo
	if err:=service.Run();err!=nil{
		panic(err)
	}

}