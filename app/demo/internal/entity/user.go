package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint64    `gorm:"id"`
	UserName  string    `gorm:"user_name"`  // 用户名
	Mobile    string    `gorm:"mobile"`     // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	Nickname  string    `gorm:"nick_name"`  // 用户昵称
	Password  string    `gorm:"password"`   // 登录密码;cmf_password加密
	Status    int8      `gorm:"status"`     // 用户状态;0:禁用,1:正常,2:未验证
	Email     string    `gorm:"email"`      // 用户登录邮箱
	Sex       int32     `gorm:"sex"`        // 性别;0:保密,1:男,2:女
	Avatar    string    `gorm:"avatar"`     // 用户头像
	Remark    string    `gorm:"remark"`     // 备注
	IsAdmin   int8      `gorm:"is_admin"`   // 是否后台管理员 1 是  0   否
	Address   string    `gorm:"address"`    // 联系地址
	CreatedAt time.Time `gorm:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"updated_at"` // 更新时间
}

// TableName 表名称
func (*User) TableName() string {
	return "tb_user"
}
