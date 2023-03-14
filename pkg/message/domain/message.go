package domain

// MessageBase 消息主表
type MessageBase struct {
	Source        string `json:"source"`
	Content       string `json:"content"`
	ReceiverRange uint   `json:"receiver_range"` // 接收者范围 1:全部 2:个人 3：某角色 4:某项目
	ReceiverId    uint   `json:"receiver_id"`
	ReadStatus    uint   `gorm:"-" json:"read_status"` // 已读状态 1:未读 2:已读
}

// MessageReadBase 已读消息表
type MessageReadBase struct {
	MessageId uint `json:"message_id"`
	UserId    uint `json:"user_id"`
}

type MessageReq struct {
	BaseModel
	MessageReadBase
}

type MessageReqPaginate struct {
	PaginateReq
	ReadStatus uint `json:"read_status"` // 已读状态 0:全部 1:未读 2:已读
	MessageScope
}

type MessageScope struct {
	Scope map[int][]string `json:"scope"` //消息范围 key:接收者范围(receiver_range) value:接收者id(receiver_id)
}

type MessageReadReq struct {
	BaseModel
	MessageReadBase
}
