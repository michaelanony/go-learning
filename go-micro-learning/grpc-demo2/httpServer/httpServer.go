package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-learning/go-micro-learning/grpc-demo2/httpServer/services"
	"go-learning/go-micro-learning/grpc-demo2/server/certs"
	"google.golang.org/grpc"
	"log"
	"net/http"
)



func main() {
	gwmux:=runtime.NewServeMux()
	opt:=[]grpc.DialOption{grpc.WithTransportCredentials(certs.GetCreds("client"))}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(),gwmux,"localhost:8081",opt)
	if err!=nil{
		log.Fatal(err)
	}
	httpServer:=&http.Server{
		Addr: ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
