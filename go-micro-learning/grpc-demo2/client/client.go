package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-learning/go-micro-learning/grpc-demo2/client/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func main() {
	//creds, _ := credentials.NewClientTLSFromFile("go-micro-learning/grpc-demo2/client/certs/server.crt", "michael.com")
	cert, _ := tls.LoadX509KeyPair("go-micro-learning/grpc-demo2/client/certs/client.pem", "go-micro-learning/grpc-demo2/client/certs/client.key")
	certPool := x509.NewCertPool()
	ca,_ :=ioutil.ReadFile("go-micro-learning/grpc-demo2/client/certs/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds :=credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName: "localhost",
		RootCAs: certPool,
	})
	conn,err:=grpc.Dial(":8081",grpc.WithTransportCredentials(creds))
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	client:=services.NewProdServiceClient(conn)
	resp, err := client.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(resp)
}


