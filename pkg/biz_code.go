package pkg

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	UnAuthorized          = errors.New(401, "401", "Unauthorized")
	CodeNotFound          = errors.New(404, "404", "Not Found")
	CodeInternalError     = errors.New(500, "500", "Internal Server Error")
	CodeUnknown           = errors.New(520, "520", "Unknown Error")
	CodeValidationFailed  = errors.New(417, "9999", "Validation Failed")
	CodeBusinessFailed    = errors.New(417, "10000", "业务出错")
	AccountPassWordFailed = errors.New(417, "10001", "账号或密码错误")
	AccountLock           = errors.New(417, "10002", "账号已冻结")
	CaptchaInvalid        = errors.New(417, "10003", "验证码输入错误")
)
