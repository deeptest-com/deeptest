package model

import (
	"gorm.io/gorm"
	"time"
)

// Oplog 中间件 modelRef
type Oplog struct {
	gorm.Model
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法" validate:"required"`
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径" validate:"required"`
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态" validate:"required"`
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟"`
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`
	Body         string        `json:"body" form:"body" gorm:"type:longtext;column:body;comment:请求Body"`
	Resp         string        `json:"resp" form:"resp" gorm:"type:longtext;column:resp;comment:响应Body"`
	UserID       uint          `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
}

func (Oplog) TableName() string {
	return "biz_oplog"
}
