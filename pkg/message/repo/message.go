package repo

import (
	"fmt"
	"gorm.io/gorm"
	"message/domain"
	"message/model"
	"strings"
)

type MessageRepo struct {
	DB              *gorm.DB         `inject:""`
	MessageReadRepo *MessageReadRepo `inject:""`
}

func (r *MessageRepo) Create(message domain.MessageReq) (id uint, err error) {
	err = r.DB.Model(&model.Message{}).Create(&message).Error
	if err != nil {
		return
	}

	id = message.Id
	return
}

func (r *MessageRepo) Paginate(req domain.MessageReqPaginate) (data domain.PageData, err error) {
	var count int64
	var messages []model.Message

	db := r.DB
	var sql, scopeSql string

	if len(req.Scope) > 0 {
		for receiverRange, receiverIds := range req.Scope {
			tmpSql := " OR (receiver_range = %d AND receiver_id = IN (%s))"
			tmpSql = fmt.Sprintf(tmpSql, receiverRange, strings.Join(receiverIds, ","))
			scopeSql = scopeSql + tmpSql
		}
	}
	//全部消息
	if req.ReadStatus == 0 {
		sql = "SELECT * FROM %s WHERE receiver_range = 1" + scopeSql
		sql = fmt.Sprintf(sql, model.Message{}.TableName())
		err = db.Count(&count).Error
		if err != nil {
			return
		}

		err = db.Scopes(PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
			Find(&messages).Error
		if err != nil {
			return
		}

		//查出列表中已读的消息
		messageIds := make([]uint, 0)
		for _, v := range messages {
			messageIds = append(messageIds, v.ID)
		}
		messagesRead, err := r.MessageReadRepo.ListByMessageIds(messageIds)

		messageReadMap := make(map[uint]uint)
		if err != nil {
			for _, v := range messagesRead {
				messageReadMap[v.MessageId] = v.MessageId
			}
		}

		for _, message := range messages {
			if _, ok := messageReadMap[message.ID]; ok {
				message.ReadStatus = 2
			} else {
				message.ReadStatus = 1
			}
		}
	} else {
		sql = "SELECT * FROM %s m LEFT JOIN %s r ON m.id=r.message_id WHERE (m.receiver_range = 1 %s ) AND r.id IS"
		//未读
		if req.ReadStatus == 1 {
			sql = sql + " NULL"
		} else if req.ReadStatus == 2 {
			//已读
			sql = sql + " NOT NULL"
		}
		sql = fmt.Sprintf(sql, model.Message{}.TableName(), model.MessageRead{}.TableName(), scopeSql)

		db = db.Raw(sql)

		err = db.Count(&count).Error
		if err != nil {
			return
		}

		err = db.Scopes(PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
			Find(&messages).Error
		if err != nil {
			return
		}

		for _, message := range messages {
			message.ReadStatus = req.ReadStatus
		}
	}

	data.Populate(messages, count, req.Page, req.PageSize)

	return
}

func (r *MessageRepo) GetUnreadCount(scope domain.MessageScope) (count int64, err error) {
	var scopeSql string

	if len(scope.Scope) > 0 {
		for receiverRange, receiverIds := range scope.Scope {
			tmpSql := " OR (receiver_range = %d AND receiver_id = IN (%s))"
			tmpSql = fmt.Sprintf(tmpSql, receiverRange, strings.Join(receiverIds, ","))
			scopeSql = scopeSql + tmpSql
		}
	}
	sql := "SELECT * FROM %s m LEFT JOIN %s r ON m.id=r.message_id WHERE (m.receiver_range = 1 %s ) AND r.id IS NULL"
	sql = fmt.Sprintf(sql, model.Message{}.TableName(), model.MessageRead{}.TableName(), scopeSql)

	err = r.DB.Raw(sql).Count(&count).Error
	if err != nil {
		return
	}
	return
}
