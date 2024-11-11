package common

import "errors"

var (
	ErrPasswordIncorrect = errors.New("密码错误")
	ErrUsernameExists    = errors.New("用户名已存在")
	ErrEmailExists       = errors.New("邮箱已存在")
	ErrPhoneExists       = errors.New("手机号已存在")
)
