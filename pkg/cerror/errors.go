package cerror

import (
	"google.golang.org/grpc/status"
)

var (
	ErrUserHasExists  = status.Error(101, "用户名已存在")
	ErrUserNotFound   = status.Error(102, "用户不存在")
	ErrEmailHasExists = status.Error(103, "邮箱已存在")
	ErrPhoneHasExists = status.Error(104, "手机号已存在")
	ErrPasswordWrong  = status.Error(105, "密码错误")
	ErrUserDisabled   = status.Error(106, "用户被禁用")
)

var (
	ErrNotLogin         = status.Error(201, "用户未登录")
	ErrPermissionDenied = status.Error(202, "权限不足")
	ErrTokenInvalid     = status.Error(203, "token无效")
	ErrTokenExpired     = status.Error(204, "token过期")
	ErrTokenNotMatch    = status.Error(205, "token不匹配")
)

var (
	ErrRoleHasExists = status.Error(301, "角色已存在")
)
