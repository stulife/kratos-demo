package model

type User struct {
	Id       uint64
	UserName string // 用户名
	Mobile   string // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	Nickname string // 用户昵称
	Password string
	Status   int8   // 用户状态;0:禁用,1:正常,2:未验证
	Email    string // 用户登录邮箱
	Sex      int32  // 性别;0:保密,1:男,2:女
	Avatar   string // 用户头像
	Remark   string // 备注
	IsAdmin  int8   // 是否后台管理员 1 是  0   否
	Address  string // 联系地址
}
