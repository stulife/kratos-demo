package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"kratos-demo/app/demo/internal/model"
)

type UserRepo interface {
	CreateUser(ctx context.Context, u *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, u *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*model.User, error)
	GetUserByName(ctx context.Context, userName string) (*model.User, error)
	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*model.User, error)
	GetRedisClient() *redis.Client
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/demo"))}
}
func (uc *UserUseCase) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
	return uc.repo.CreateUser(ctx, u)
}
func (uc *UserUseCase) UpdateUser(ctx context.Context, u *model.User) (*model.User, error) {
	return uc.repo.UpdateUser(ctx, u)
}
func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) {
	uc.repo.DeleteUser(ctx, id)
}
func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return uc.repo.GetUser(ctx, id)
}
func (uc *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*model.User, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) GetUserByName(ctx context.Context, userName string) (*model.User, error) {
	return uc.repo.GetUserByName(ctx, userName)
}
