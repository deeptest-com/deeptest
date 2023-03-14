package domain

import _domain "github.com/aaronchen2k/deeptest/pkg/domain"

type MessageReq struct {
	_domain.Model
	MessageBase
}

type MessageReqPaginate struct {
	_domain.PaginateReq

	ReadStatus uint `json:"read_status"` // 已读状态 0:全部 1:未读 2:已读
}

type MessageBase struct {
	Source        string `json:"source"`
	Content       string `json:"content"`
	ReceiverRange uint   `json:"receiver_range"` // 接收者范围 1:全部 2:个人 3：某角色 4:某项目
	ReceiverId    uint   `json:"receiver_id"`
	ReadStatus    uint   `gorm:"-" json:"read_status"` // 已读状态 1:未读 2:已读
}

type MessageResp struct {
	_domain.PaginateReq
	MessageBase
}

type MessageReadReq struct {
	_domain.Model
	MessageReadBase
}

type MessageReadBase struct {
	MessageId uint `json:"message_id"`
	UserId    uint `json:"user_id"`
}
