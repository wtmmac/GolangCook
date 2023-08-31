package service

import (
	"context"
	"time"
)
import pb "github.com/wtmmac/GolangCook/grpc/api"

type HelloService struct{}

func (s *HelloService) SayHello(c context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	ret := new(pb.HelloReply)
	ret.Message = "hello, " + r.Name + " " + time.Now().String()
	return ret, nil
}
