package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"kratos-demo/app/demo/internal/biz"
	"kratos-demo/app/demo/internal/entity"
	"kratos-demo/app/demo/internal/model"
	"time"
)

//var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/demo")),
	}
}

func (r userRepo) CreateUser(ctx context.Context, b *model.User) (*model.User, error) {
	u := entity.User{
		UserName:  b.UserName,
		Mobile:    b.Mobile,
		Nickname:  b.Nickname,
		Password:  b.Password,
		Email:     b.Email,
		Sex:       b.Sex,
		Avatar:    b.Avatar,
		Remark:    b.Remark,
		Address:   b.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := r.data.db.WithContext(ctx).Create(&u)
	return &model.User{Id: u.ID}, result.Error
}

func (r userRepo) UpdateUser(ctx context.Context, u *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepo) DeleteUser(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r userRepo) GetUser(ctx context.Context, id int64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}
func (r userRepo) GetUserByName(ctx context.Context, userName string) (*model.User, error) {

	u := entity.User{}

	result := r.data.db.WithContext(ctx).Where(&entity.User{UserName: userName}).First(&u)

	return &model.User{
		Id:       u.ID,
		UserName: u.UserName,
		Mobile:   u.Mobile,
		Nickname: u.Nickname,
		Password: u.Password,
		Status:   u.Status,
		Email:    u.Email,
		Sex:      u.Sex,
		Avatar:   u.Avatar,
		Remark:   u.Remark,
		IsAdmin:  u.IsAdmin,
		Address:  u.Address,
	}, result.Error
}

func (r userRepo) GetRedisClient() *redis.Client {
	return r.data.rdb
}
