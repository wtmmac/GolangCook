package main

import (
	"context"
	"fmt"
	pb "github.com/wtmmac/GolangCook/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func myClient() {
	fmt.Printf("%c[0;47;42m%s%c[0m\n", 0x1B, ">>>>>>>>>>>>>"+time.Now().String(), 0x1B)
	conn, err := grpc.Dial("tonymac:3399", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// handle error
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)

	for i := 0; i < 3; i++ {
		req := pb.HelloRequest{
			Name: "world",
		}
		reply, err := client.SayHello(context.Background(), &req)
		if err != nil {
			fmt.Println("client.SayHello error:", err)
			return
		}

		fmt.Printf("get msg from server:[%v] \n", reply)
	}
}

func main() {
	myClient()
}
