package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type WsReq struct {
	Act consts.ExecType `json:"act"`
	Id  int             `json:"id"`
}
