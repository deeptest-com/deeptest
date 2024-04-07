package serverDomain

type RecordReq struct {
	TargetId uint         `json:"targetId"`
	ServeId  uint         `json:"serveId"`
	UserId   uint         `json:"userId"`
	Items    []RecordItem `json:"items"`
}

type RecordItem struct {
	Request  RecordRequest  `json:"request"`
	Response RecordResponse `json:"response"`
}
type RecordResponse struct {
	Body       string `json:"body"`
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

type RecordRequest struct {
	HasPostData      bool                              `json:"hasPostData"`
	Headers          map[string]string                 `json:"headers"`
	Cookies          map[string]map[string]interface{} `json:"cookies"`
	InitialPriority  string                            `json:"initialPriority"`
	IsSameSite       bool                              `json:"isSameSite"`
	Method           string                            `json:"method"`
	MixedContentType string                            `json:"mixedContentType"`
	PostData         string                            `json:"postData"`
	Body             string                            `json:"body"`
	ReferrerPolicy   string                            `json:"referrerPolicy"`
	Url              string                            `json:"url"`

	//PostDataEntries  []struct {
	//	Bytes string `json:"bytes"`
	//} `json:"postDataEntries"`
}
