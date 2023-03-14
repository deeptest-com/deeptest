package domain

type BaseModel struct {
	Id        uint   `json:"id"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

type PaginateReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Field    string `json:"field"`
	Order    string `json:"order"`
}

type PageData struct {
	Result interface{} `json:"result"`

	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (d *PageData) Populate(result interface{}, total int64, page, pageSize int) {
	d.Result = result
	d.Total = int(total)
	d.Page = page
	d.PageSize = pageSize
}