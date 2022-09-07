package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"kratos-demo/pkg"
	"kratos-demo/pkg/util"
	"time"
)

var currentUserKey struct{}

type CurrentUser struct {
	Id         uint64
	UserName   string // 用户名
	Nickname   string
	Mobile     string
	Email      string
	uuid       string //token key
	loginTime  int64  // 登录时间
	expireTime int64  // 过期时间
}

func GenerateToken(ctx context.Context, secret string, duration time.Duration, rdb *redis.Client, u *CurrentUser) string {
	u.uuid = uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": u.uuid,
		//"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	refreshToken(ctx, duration, rdb, u)
	return tokenString
}

func JWTAuth(header string, secret string, duration time.Duration, rdb *redis.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				tokenString := tr.RequestHeader().Get(header)
				//tokenString := tr.RequestHeader().Get("Authorization")
				//auths := strings.SplitN(tokenString, " ", 2)
				//if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
				//	return nil, pkg.UnAuthorized
				//}

				if tokenString == "" {
					return nil, pkg.UnAuthorized
				}
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(secret), nil
				})

				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					// put CurrentUser into ctx
					if u, ok := claims["uuid"]; ok {
						var currentUser = fromRedis(ctx, rdb, u.(string))
						if currentUser == nil {
							return nil, pkg.UnAuthorized
						}
						currentUser.uuid = u.(string)
						verifyToken(ctx, duration, rdb, currentUser)
						ctx = WithContext(ctx, currentUser)
					}
				} else {
					return nil, pkg.UnAuthorized
				}
			}
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) *CurrentUser {
	return ctx.Value(currentUserKey).(*CurrentUser)
}

func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}

func refreshToken(ctx context.Context, duration time.Duration, rdb *redis.Client, u *CurrentUser) {
	u.loginTime = time.Now().UnixMilli()
	u.expireTime = u.loginTime + duration.Milliseconds()
	userKey := util.LoginKey(u.uuid)
	buf, _ := json.Marshal(u)
	rdb.Set(ctx, userKey, buf, duration)

}

func verifyToken(ctx context.Context, duration time.Duration, rdb *redis.Client, u *CurrentUser) {
	var expireTime = u.expireTime
	var currentTime = time.Now().UnixMilli()
	if expireTime-currentTime <= (20 * (time.Minute)).Milliseconds() {
		refreshToken(ctx, duration, rdb, u)
	}
}

func fromRedis(ctx context.Context, rdb *redis.Client, uuid string) *CurrentUser {
	userKey := util.LoginKey(uuid)
	var currentUser *CurrentUser
	buf, _ := rdb.Get(ctx, userKey).Bytes()
	json.Unmarshal(buf, &currentUser)
	return currentUser
}
