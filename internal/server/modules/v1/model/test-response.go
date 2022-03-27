package model

type TestResponse struct {
	BaseModel

	HttpCode string `json:"httpCode"`
}

func (TestResponse) TableName() string {
	return "biz_test_response"
}
