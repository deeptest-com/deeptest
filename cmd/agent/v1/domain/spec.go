package domain

type ParseSpecReq struct {
	Src  string `json:"src"`
	Type string `json:"type"`
	File string `json:"file"`
	Url  string `json:"url"`

	TargetId  int    `json:"targetId"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}

type SubmitSpecResp struct {
}
