package message

import "github.com/aaronchen2k/deeptest/pkg/message/domain"

type Handler interface {
	Create(req domain.MessageReq) (uint, error)
	Paginate(req domain.MessageReqPaginate) (ret domain.PageData, err error)
	UnreadCount(scope domain.MessageScope) (count int64, err error)
	OperateRead(req domain.MessageReadReq) (uint, error)
}
