package agentExec

import (
	"time"
)

type ProcessorCookie struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	CookieName   string     `json:"cookieName" yaml:"cookieName"`
	VariableName string     `json:"variableName" yaml:"variableName"`
	RightValue   string     `json:"rightValue" yaml:"rightValue"`
	Domain       string     `json:"domain" yaml:"domain"`
	ExpireTime   *time.Time `json:"expireTime" yaml:"expireTime"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}

func (p ProcessorCookie) Run(s *Session) (log Log, err error) {
	return
}
