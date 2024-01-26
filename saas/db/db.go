package db

type Config struct {
	Path            string `json:"path"`
	Config          string `json:"Config"`
	Dbname          string `json:"dbname"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	SchemaType      string `json:"schemaType"`
	Maxidleconns    string `json:"maxidleconns"`
	Maxopenconns    string `json:"maxopenconns"`
	Connmaxlifetime string `json:"connmaxlifetime"`
}
