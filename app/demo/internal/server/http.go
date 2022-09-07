package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	v1 "kratos-demo/api/demo/v1"
	"kratos-demo/app/demo/internal/conf"
	"kratos-demo/app/demo/internal/service"
	validate "kratos-demo/pkg/middleware"
	"kratos-demo/pkg/middleware/auth"
	"kratos-demo/pkg/util"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, s *service.UserService, l *service.LoginService, jwtc *conf.Jwt, rdb *redis.Client, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(auth.JWTAuth(jwtc.Header, jwtc.Secret, jwtc.ExpireTime.AsDuration(), rdb)).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.ResponseEncoder(util.CustomResponseEncoder),
		http.ErrorEncoder(util.CustomErrorEncoder),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "PATCH", "HEAD", "CONNECT", "OPTIONS", "TRACE"}),
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "User-Agent", "Cookie", "Authorization", "X-Auth-Token", "X-Requested-With"}),
			handlers.MaxAge(3628800),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIHandler)

	v1.RegisterUserHTTPServer(srv, s)
	v1.RegisterLoginHTTPServer(srv, l)

	return srv
}

func NewSkipRoutersMatcher() selector.MatchFunc {
	skipRouters := make(map[string]struct{})
	skipRouters["/api.demo.v1.Login/Login"] = struct{}{}
	skipRouters["/api.demo.v1.Login/GetCaptcha"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}
