package model

type EndpointTagRel struct {
	BaseModel
	EndpointId uint `json:"endpointId"`
	TagId      uint `json:"tagId"`
}

func (EndpointTagRel) TableName() string {
	return "biz_endpoint_tag_rel"
}
