package model

type EndpointFavorite struct {
	BaseModel
	UserId     uint
	EndpointId uint
}

func (EndpointFavorite) TableName() string {
	return "biz_endpoint_favorite"
}
