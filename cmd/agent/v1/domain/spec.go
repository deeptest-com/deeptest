package domain

type SubmitSpecReq struct {
	Src     string `json:"src"`
	Type    string `json:"type"`
	Content string `json:"content"`
	File    string `json:"file"`
	Url     string `json:"url"`
}

type SubmitSpecResp struct {
}
