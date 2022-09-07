package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	v1 "kratos-demo/api/demo/v1"
	"kratos-demo/app/demo/internal/conf"
	"kratos-demo/app/demo/internal/model"
	"kratos-demo/pkg"
	"kratos-demo/pkg/crypto"
	"kratos-demo/pkg/middleware/auth"
	"strings"
)

type LoginRepo interface {
}

type LoginUseCase struct {
	repo   UserRepo
	driver *base64Captcha.DriverString
	store  base64Captcha.Store
	jwtc   *conf.Jwt
	log    *log.Helper
}

func NewLoginUseCase(repo UserRepo, logger log.Logger, jwtc *conf.Jwt) *LoginUseCase {
	return &LoginUseCase{repo: repo, driver: &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      50,
		ShowLineOptions: 20,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:           []string{"chromohv.ttf"},
	}, store: base64Captcha.DefaultMemStore, jwtc: jwtc, log: log.NewHelper(log.With(logger, "module", "usecase/demo"))}
}

func (s *LoginUseCase) GetCaptcha() (id, b64s string, err error) {
	driver := s.driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, s.store)
	id, b64s, err = c.Generate()
	return

}

func (s *LoginUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {

	//if !s.verifyCaptcha(req.VerifyKey, req.VerifyCode) {
	//	return nil, pkg.CaptchaInvalid
	//}
	u, _ := s.repo.GetUserByName(ctx, req.Username)

	if !crypto.CheckPasswordHash(req.Password, u.Password) {

		return nil, pkg.AccountPassWordFailed
	}
	uInfo := &v1.UserRes{
		Id:       u.Id,
		Nickname: u.Nickname,
		Mobile:   u.Mobile,
		Email:    u.Email,
		Username: u.UserName,
	}
	return &v1.LoginReply{Token: s.generateToken(ctx, s.repo.GetRedisClient(), u), UserInfo: uInfo}, nil
}

func (s *LoginUseCase) verifyCaptcha(id, answer string) bool {
	c := base64Captcha.NewCaptcha(s.driver, s.store)
	answer = strings.ToLower(answer)
	return c.Verify(id, answer, true)

}
func (s *LoginUseCase) generateToken(ctx context.Context, rdb *redis.Client, u *model.User) string {

	currentUser := &auth.CurrentUser{
		Id:       u.Id,
		UserName: u.UserName,
		Mobile:   u.Mobile,
		Nickname: u.Nickname,
		Email:    u.Email,
	}
	return auth.GenerateToken(ctx, s.jwtc.Secret, s.jwtc.ExpireTime.AsDuration(), rdb, currentUser)
}
