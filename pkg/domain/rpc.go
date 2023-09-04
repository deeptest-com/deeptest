package _domain

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/pkg/consts"
)

type RpcReq struct {
	NodeIp   string
	NodePort int

	ApiPath   string
	ApiMethod string
	Data      interface{}
}

type RpcResp struct {
	Code    _consts.ResultCode `json:"code"`
	Msg     string             `json:"msg"`
	Payload interface{}        `json:"payload"`
}

func (result *RpcResp) Pass(msg string) {
	result.Code = _consts.ResultCodeSuccess
	result.Msg = msg
}

func (result *RpcResp) Passf(str string, args ...interface{}) {
	result.Code = _consts.ResultCodeSuccess
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RpcResp) Fail(msg string) {
	result.Code = _consts.ResultCodeFail
	result.Msg = msg
}

func (result *RpcResp) Failf(str string, args ...interface{}) {
	result.Code = _consts.ResultCodeFail
	result.Msg = fmt.Sprintf(str, args...)
}

func (result *RpcResp) IsSuccess() bool {
	return result.Code == _consts.ResultCodeSuccess
}
