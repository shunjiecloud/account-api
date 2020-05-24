package ec

import (
	"net/http"

	ec "github.com/shunjiecloud/pkg/errcode"
)

var (
	ErrEmailAlreadyUsed    = ec.Error{Result: "ErrEmailAlreadyUsed", Msg: "用户邮箱已被使用", HttpCode: http.StatusBadRequest, IsWarn: true}
	ErrUserNameAlreadyUsed = ec.Error{Result: "ErrUserNameAlreadyUsed", Msg: "用户名已被使用", HttpCode: http.StatusBadRequest, IsWarn: true}
)
