package domain

type DbConfig struct {
	Path            string `json:"path"`
	Config          string `json:"Config"`
	Dbname          string `json:"dbname"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	SchemaType      string `json:"schemaType"`
	Maxidleconns    int64  `json:"maxidleconns"`
	Maxopenconns    int64  `json:"maxopenconns"`
	Connmaxlifetime int64  `json:"connmaxlifetime"`
}

type Tenant struct {
	Id       int64    `json:"id"`
	DbConfig DbConfig `json:"pjtDB"`
}
