package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-demo/api/demo/v1"
	"kratos-demo/app/demo/internal/biz"
)

type LoginService struct {
	v1.UnimplementedUserServer

	oc  *biz.LoginUseCase
	log *log.Helper
}

func NewLoginService(oc *biz.LoginUseCase, logger log.Logger) *LoginService {
	return &LoginService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/demo"))}
}

func (s *LoginService) GetCaptcha(ctx context.Context, req *v1.GetCaptchaRequest) (res *v1.GetCaptchaReply, err error) {

	var (
		id, b64s string
	)
	id, b64s, err = s.oc.GetCaptcha()

	return &v1.GetCaptchaReply{
		Id:           id,
		Base64String: b64s,
	}, err
}
func (s *LoginService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.oc.Login(ctx, req)
}
