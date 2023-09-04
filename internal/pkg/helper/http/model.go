package httpHelper

type Request struct {
	PageSize int `json:"pageSize"`
	PageNo   int `json:"pageNo"`
	Aa       int `json:"aa"`
}

type Response struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message interface{} `json:"message"`
	Result  interface{} `json:"result"`
}
