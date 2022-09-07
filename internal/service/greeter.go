package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	v1 "kratos-demo/api/helloworld/v1"
	"kratos-demo/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	//panic("dfadsf")
	//return nil, errors.New(401, "Authentication failed", "Missing token or token incorrect")
	return nil, errors.New(400, "sdfsd", "sdf")
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
