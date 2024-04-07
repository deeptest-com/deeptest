package model

type EndpointFavorite struct {
	BaseModel
	UserId     uint `json:"userId"`
	EndpointId uint `json:"endpointId"`
}

func (EndpointFavorite) TableName() string {
	return "biz_endpoint_favorite"
}
