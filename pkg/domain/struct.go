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

type ErrMsg struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
}

var (
	NoErr         = ErrMsg{0, "请求成功"}
	NeedInitErr   = ErrMsg{2001, "前往初始化数据库"}
	AuthErr       = ErrMsg{4001, "请重新登录"}
	AuthActionErr = ErrMsg{4003, "权限错误"}
	ParamErr      = ErrMsg{4004, "参数解析失败"}
	SystemErr     = ErrMsg{5000, "系统错误"}

	BizErrNameExist = ErrMsg{10100, "指定名称的记录不存在"}
)
