package serverDomain

import "github.com/deeptest-com/deeptest/pkg/consts"

type RpcResult struct {
	Code    _consts.ResultCode `json:"code"`
	Msg     string             `json:"msg"`
	Payload interface{}        `json:"payload"`
}

func (result *RpcResult) Pass(msg string) {
	result.Code = _consts.ResultCodeSuccess
	result.Msg = msg
}

func (result *RpcResult) Fail(msg string) {
	result.Code = _consts.ResultCodeFail
	result.Msg = msg
}

func (result *RpcResult) IsSuccess() bool {
	return result.Code == _consts.ResultCodeSuccess
}
