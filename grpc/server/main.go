package main

import (
	"fmt"
	pb "github.com/wtmmac/GolangTest/grpc/api"
	"github.com/wtmmac/GolangTest/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

func myServer() {
	fmt.Printf("%c[0;47;42m%s%c[0m\n", 0x1B, ">>>>>>>>>>>>>"+time.Now().String(), 0x1B)
	lis, err := net.Listen("tcp", ":3399")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, &service.HelloService{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

func main() {
	myServer()
}
