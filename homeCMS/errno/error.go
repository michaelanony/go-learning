package errno

import "errors"

var(
	ERROR_USER_NOT_EXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS = errors.New("用户已经存在")
	ERROR_USER_PWD = errors.New("用户密码错误")
	ERROR_USER_FORMAT = errors.New("用户信息格式错误")
)
