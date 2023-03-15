package message

import (
	"github.com/aaronchen2k/deeptest/pkg/message/domain"
	"gorm.io/gorm"
)

type Client struct {
	Handler Handler    `inject:""`
	Redis   *RedisConf `json:"redis"`
	Mysql   *MysqlConf `json:"mysql"`
	Db      *gorm.DB
}

func NewMessageService() *Client {
	return &Client{}
}

func init() {
	client := NewMessageService()
	var err error
	if err = client.CheckRedis(); err != nil {
		return
	}
	if err = client.InitMysql(); err != nil {
		return
	}
	if err = client.CheckMessageTable(); err != nil {
		return
	}
}

func (c *Client) Create(req domain.MessageReq) (uint, error) {
	return c.Handler.Create(req)
}

func (c *Client) Paginate(req domain.MessageReqPaginate) (ret domain.PageData, err error) {
	return c.Handler.Paginate(req)
}

func (c *Client) UnreadCount(req domain.MessageScope) (count int64, err error) {
	return c.Handler.UnreadCount(req)
}

func (c *Client) OperateRead(req domain.MessageReadReq) (uint, error) {
	return c.Handler.OperateRead(req)
}
