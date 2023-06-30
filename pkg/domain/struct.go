package _domain

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

type Model struct {
	Id        uint   `json:"id"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

type ReqId struct {
	Id uint `json:"id" param:"id"`
}

type PaginateReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Field    string `json:"field"`
	Order    string `json:"order"`
}

func (r *PaginateReq) ConvertParams() {
	r.Field = stringUtils.SnakeCase(r.Field)
	r.Order = serverConsts.SortMap[r.Order]
}

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type PageData struct {
	Result interface{} `json:"result"`

	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (d *PageData) Populate(result interface{}, total int64, page, pageSize int) {
	d.Result = result
	d.Total = int(total)
	d.Page = page
	d.PageSize = pageSize
}

type BizErr struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
}

var (
	NoErr         = BizErr{0, "请求成功"}
	AuthErr       = BizErr{401, "请重新登录"}
	AuthActionErr = BizErr{403, "权限不足"}

	NeedInitErr = BizErr{1000, "未初始化"}
	ParamErr    = BizErr{2000, "参数解析失败"}

	RequestErr = BizErr{3000, "请求失败"}
	FailErr    = BizErr{3500, "处理失败"}
	SystemErr  = BizErr{4000, "系统错误"}
	LoginErr   = BizErr{5000, "登录失败"}

	ErrNoUser             = BizErr{10100, "找不到用户"}
	ErrNameExist          = BizErr{10100, "同名记录已存在"}
	ErrUsernameExist      = BizErr{10200, "用户名已占用"}
	ErrEmailExist         = BizErr{10300, "邮箱已存在"}
	ErrPasswordMustBeSame = BizErr{10500, "两次密码必须一样"}
)

func (e BizErr) Error() string {
	return e.Msg
}
