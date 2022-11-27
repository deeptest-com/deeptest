package domain

import _domain "github.com/aaronchen2k/deeptest/pkg/domain"

// LoginReq 登录请求字段
type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	_domain.ReqId
	Password string `json:"password"`
}

type RegisterReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Confirm  string `json:"confirm" validate:"required"`
}

type ResetPasswordReq struct {
	Username string `json:"username" validate:"required"`
	Vcode    string `json:"vcode" validate:"required"`
	Password string `json:"password" validate:"required"`
	Confirm  string `json:"confirm" validate:"required"`
}
