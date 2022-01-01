package serverDomain

// LoginReq 登录请求字段
type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
