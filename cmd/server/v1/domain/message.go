package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type MessageReq struct {
	_domain.Model
	MessageBase
}

type MessageReqPaginate struct {
	_domain.PaginateReq

	ReadStatus uint `json:"read_status"` // 已读状态 0:全部 1:未读 2:已读
	MessageScope
}

type MessageScope struct {
	Scope map[int][]string `json:"scope"` //消息范围 key:接收者范围(receiver_range) value:接收者id(receiver_id)
}

type MessageBase struct {
	MessageSource consts.MessageSource      `json:"message_source"` //业务模块
	Content       string                    `gorm:"type:text" json:"content"`
	ReceiverRange uint                      `json:"receiver_range"` // 接收者范围 1:全部 2:个人 3：某角色 4:某项目
	SenderId      uint                      `json:"sender_id"`      //消息发送者
	ReceiverId    uint                      `json:"receiver_id"`
	ReadStatus    uint                      `gorm:"-" json:"read_status"` // 已读状态 1:未读 2:已读
	SendStatus    consts.MessageSendStatus  `json:"send_status"`
	ServiceType   consts.MessageServiceType `json:"service_type"`
	McsMessageId  string                    `json:"mcs_message_id"`
	BusinessId    uint                      `json:"business_id"` // 业务ID
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

type McsApprovalResData struct {
	InstanceId  string   `json:"instanceId"`  //审批id
	Status      int      `json:"status"`      //审批状态 1拒绝 2同意 3终止 4取消 5进行中(企微审批创建成功会立即返回此状态)
	ApplyTime   int      `json:"applyTime"`   //发起时间：时间戳
	Comments    string   `json:"comments"`    //审批备注
	ProcessTime string   `json:"processTime"` //处理时间
	ApproveUser []string `json:"approveUser"` //审批人第三方账号
}

type McsApprovalRes struct {
	Data string `json:"data"`
}
