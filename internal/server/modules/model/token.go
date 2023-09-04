package model

type Auth2Token struct {
	BaseModel

	Name      string `json:"name"`
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
	ProjectId int    `json:"projectId"`
}

func (Auth2Token) TableName() string {
	return "biz_auth2_token"
}
