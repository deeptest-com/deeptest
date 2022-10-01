package domain

type DataReq struct {
	ClearData bool    `json:"clearData"`
	Sys       DataSys `json:"sys"`
}

type DataSys struct {
	AdminPassword string `json:"adminPassword"`
}

type DataCache struct {
	Host     string `json:"host"  validate:"required"`
	Port     string `json:"port"  validate:"required"`
	Password string `json:"password"`
	PoolSize int    `json:"poolSize"`
	DB       int    `json:"db"`
}
