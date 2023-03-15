package model

type UserList struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type SummaryDetails struct {
	BaseModel

	ProjectId          int64   `json:"project_id"`
	ProjectName        string  `gorm:"type:text" json:"project_name"`
	ProjectChineseName string  `gorm:"type:text" json:"project_chinese_name"`
	ScenarioTotal      int64   `json:"scenario_total"`
	InterfaceTotal     int64   `json:"interface_total"`
	ExecTotal          int64   `json:"exec_total"`
	PassRate           float64 `json:"pass_rate"`
	Coverage           float64 `json:"coverage"`
	AdminUser          string  `gorm:"type:text" json:"admin_user"`
	ProjectCreateTime  string  `gorm:"type:text" json:"project_create_time"`
	ProjectDesc        string  `gorm:"type:text" json:"project_desc"`
	UserList           string  `gorm:"type:text;" json:"user_list"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

type SummaryCard struct {
	SummaryCardTotal
	ProjectTotal int64   `json:"project_total"`
	InterfaceHB  float64 `json:"interface_hb"`
	ScenarioHB   float64 `json:"scenario_hb"`
	CoverageHB   float64 `json:"coverage_hb"`
}

type SummaryCardTotal struct {
	ScenarioTotal  int64               `gorm:"column:scenario_total" json:"scenario_total"`
	InterfaceTotal int64               `gorm:"column:interface_total" json:"interface_total"`
	ExecTotal      int64               `gorm:"column:exec_total" json:"exec_total"`
	PassRate       float64             `gorm:"column:pass_rate" json:"pass_rate"`
	Coverage       float64             `gorm:"column:coverage" json:"coverage"`
	Logs           []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (SummaryDetails) TableName() string {
	return "biz_summary_details"
}
