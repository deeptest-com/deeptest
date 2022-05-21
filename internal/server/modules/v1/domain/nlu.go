package serverDomain

import (
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
	"time"
)

type NluTask struct {
	Version string      `yaml:"version" default:"2.0"`
	Name    string      `yaml:"name"`
	Intents []NluIntent `yaml:"intents,flow"`
}

type NluIntent struct {
	Id   uint   `yaml:"id"`
	Name string `yaml:"name"`

	Sents []NluSent `yaml:"examples"`
}

type NluSent struct {
	Id      uint     `yaml:"id"`
	Example string   `yaml:"example"`
	Slots   []string `yaml:"slots"`
}

type NluReq struct {
	AgentId    int    `json:"agentId"`
	Text       string `json:"text"`
	TextOrigin string `json:"textOrigin"`
}

type NluResp struct {
	Text string             `json:"text"`
	Code _consts.ResultCode `json:"code"`

	RasaResult *NluResult         `json:"nluResult,omitempty"`
	ExecResult *InstructionResult `json:"execResult,omitempty"`
	Msg        *map[string]string `json:"msg,omitempty"`
}

func (resp *NluResp) SetResult(result NluResult) {
	resp.RasaResult = &result
}

func (resp *NluResp) SetMsg(msg map[string]string) {
	resp.Msg = &msg
}

type NluResult struct {
	ResponseSelector *ResponseSelector `json:"response_selector,omitempty"`
	TextOrigin       string            `json:"textOrigin,omitempty"`

	Entities      []Entity        `json:"entities"`
	Intent        *Intent         `json:"intent"`
	IntentRanking []IntentRanking `json:"intent_ranking"`
	Text          string          `json:"text"`

	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

type Entity struct {
	ConfidenceEntity float64  `json:"confidence_entity"`
	End              int64    `json:"end"`
	Entity           string   `json:"entity"`
	Extractor        string   `json:"extractor"`
	Processors       []string `json:"processors"`
	Start            int64    `json:"start"`
	Value            string   `json:"value"`
	ValueOrigin      string   `json:"valueOrigin,omitempty"`
}

type Intent struct {
	Confidence float32 `json:"confidence"`
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Sent       Sent    `json:"sent,omitempty"`
}
type Sent struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type IntentRanking struct {
	Confidence int64  `json:"confidence"`
	ID         int64  `json:"id"`
	Name       string `json:"name"`
}
type ResponseSelector struct {
	AllRetrievalIntents []interface{} `json:"all_retrieval_intents"`
	Default             Default       `json:"default"`
}
type Default struct {
	Ranking  []interface{} `json:"ranking"`
	Response Response      `json:"response"`
}
type Response struct {
	Confidence        int64       `json:"confidence"`
	ID                interface{} `json:"id"`
	IntentResponseKey interface{} `json:"intent_response_key"`
	ResponseTemplates interface{} `json:"response_templates"`
	Responses         interface{} `json:"responses"`
	TemplateName      string      `json:"template_name"`
	UtterAction       string      `json:"utter_action"`
}

type InstructionResult struct {
	Code    _consts.ResultCode `json:"code"`
	Msg     string             `json:"msg"`
	Payload interface{}        `json:"payload"`

	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

func (result *InstructionResult) Pass(msg string) {
	result.Code = _consts.ResultCodeSuccess

	if msg == "" {
		msg = "success"
	}
	result.Msg = msg
}

func (result *InstructionResult) Fail(msg string) {
	result.Code = _consts.ResultCodeFail
	result.Msg = msg
}

func (result *InstructionResult) IsSuccess() bool {
	return result.Code == _consts.ResultCodeSuccess
}
