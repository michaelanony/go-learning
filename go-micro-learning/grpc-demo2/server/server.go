package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-learning/go-micro-learning/grpc-demo2/server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net/http"
)

func main()  {
	//creds, err := credentials.NewServerTLSFromFile("go-micro-learning/grpc-demo2/server/certs/server.crt", "go-micro-learning/grpc-demo2/server/certs/server_no_passwd.key")
	//if err!=nil{
	//	log.Fatal(err)
	//}
	cert, _ := tls.LoadX509KeyPair("go-micro-learning/grpc-demo2/server/certs/server.pem", "go-micro-learning/grpc-demo2/server/certs/server.key")
	certPool := x509.NewCertPool()
	ca,_ :=ioutil.ReadFile("go-micro-learning/grpc-demo2/server/certs/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds :=credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	})
	rpcServer:=grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	//lis,_:=net.Listen("tcp",":8081")
	//rpcServer.Serve(lis)
	//使用http服务
	mux:=http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Proto)
		rpcServer.ServeHTTP(writer,request)
	})
	httpServer:=&http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	httpServer.ListenAndServeTLS("go-micro-learning/grpc-demo2/server/certs/server.crt", "go-micro-learning/grpc-demo2/server/certs/server_no_passwd.key")
}
