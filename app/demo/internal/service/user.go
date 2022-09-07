package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-demo/api/demo/v1"
	"kratos-demo/app/demo/internal/biz"
	"kratos-demo/app/demo/internal/model"
	"kratos-demo/pkg/crypto"
	"kratos-demo/pkg/middleware/auth"
)

type UserService struct {
	v1.UnimplementedUserServer

	oc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(oc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/demo"))}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.UserReply, error) {

	user, err := s.oc.CreateUser(ctx, &model.User{

		UserName: req.Username,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: crypto.HashPassword(req.Password),
		Email:    req.Email,
		Sex:      req.Sex,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{Id: user.Id,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Username: user.UserName}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {

	return &v1.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {

	s.oc.DeleteUser(ctx, 1)
	return &v1.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.UserReply, error) {

	currentUser := auth.FromContext(ctx)
	return &v1.UserReply{
		Id:       currentUser.Id,
		Nickname: currentUser.Nickname,
		Mobile:   currentUser.Mobile,
		Email:    currentUser.Email,
		Username: currentUser.UserName,
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return &v1.ListUserReply{}, nil
}
