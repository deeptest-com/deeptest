package domain

type SchemaReq struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type SecuritySchemaReq struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Key    string `json:"key"`
	in     string `json:"in"`
	scheme string `json:"scheme"`
	flows  string `json:“flows”`
	Title  string `json:"title"`
}
