package certs

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func GetCreds(name string) (cred credentials.TransportCredentials) {
	if name == "client"{
		cert, _ := tls.LoadX509KeyPair("go-micro-learning/grpc-demo2/client/certs/client.pem", "go-micro-learning/grpc-demo2/client/certs/client.key")
		certPool := x509.NewCertPool()
		ca,_ :=ioutil.ReadFile("go-micro-learning/grpc-demo2/client/certs/ca.pem")
		certPool.AppendCertsFromPEM(ca)
		cred =credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},
			ServerName: "localhost",
			RootCAs: certPool,
		})
		return
	}else {
		cert, _ := tls.LoadX509KeyPair("go-micro-learning/grpc-demo2/server/certs/server.pem", "go-micro-learning/grpc-demo2/server/certs/server.key")
		certPool := x509.NewCertPool()
		ca,_ :=ioutil.ReadFile("go-micro-learning/grpc-demo2/server/certs/ca.pem")
		certPool.AppendCertsFromPEM(ca)
		cred = credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs: certPool,
		})
		return
	}
}
