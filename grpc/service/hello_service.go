package service

import (
	"context"
	"github.com/wtmmac/GolangCook/grpc/api"
	"time"
)

type HelloService struct{}

func (s *HelloService) mustEmbedUnimplementedHelloServer() {
	panic("implement me")
}

func (s *HelloService) SayHello(c context.Context, r *api.HelloRequest) (*api.HelloReply, error) {
	ret := new(api.HelloReply)
	ret.Message = "hello, " + r.Name + " " + time.Now().String()
	return ret, nil
}
