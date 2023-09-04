package _domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type WsResp struct {
	Msg       string               `json:"msg"`
	IsRunning string               `json:"isRunning,omitempty"`
	Category  consts.WsMsgCategory `json:"category"`

	Info iris.Map    `json:"info,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type MqMsg struct {
	Namespace string `json:"namespace"`
	Room      string `json:"room"`
	Event     string `json:"event"`
	Content   string `json:"content"`
}
