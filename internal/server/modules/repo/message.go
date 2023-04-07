package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type MessageRepo struct {
	DB              *gorm.DB         `inject:""`
	BaseRepo        *BaseRepo        `inject:""`
	ProjectRepo     *ProjectRepo     `inject:""`
	MessageReadRepo *MessageReadRepo `inject:""`
}

func NewMessageRepo() *MessageRepo {
	return &MessageRepo{}
}

func (r *MessageRepo) GetScope(userId uint) (scope map[int][]string) {
	scope = make(map[int][]string)
	scope[2] = []string{strconv.Itoa(int(userId))}

	// 获取用户关联的项目和角色
	userProjectIds, userRoleIds := r.ProjectRepo.GetProjectsAndRolesByUser(userId)

	var userRoleIdArr, userProjectIdArr []string

	for _, v := range userRoleIds {
		userRoleIdArr = append(userRoleIdArr, strconv.Itoa(int(v)))
	}
	scope[3] = userRoleIdArr

	for _, v := range userProjectIds {
		userProjectIdArr = append(userProjectIdArr, strconv.Itoa(int(v)))
	}
	scope[4] = userProjectIdArr

	return
}

func (r *MessageRepo) Create(req v1.MessageReq) (id uint, bizErr *_domain.BizErr) {
	message := model.Message{MessageBase: req.MessageBase}
	err := r.DB.Model(&model.Message{}).Create(&message).Error
	if err != nil {
		logUtils.Errorf("add message error", zap.String("error:", err.Error()))
		bizErr.Code = _domain.SystemErr.Code
		return
	}

	id = message.ID
	return
}

func (r *MessageRepo) Paginate(req v1.MessageReqPaginate) (data _domain.PageData, err error) {
	var count int64
	var messages []model.Message

	db := r.DB
	var selectSql, countSql, sql, scopeSql string

	//全部消息
	if req.ReadStatus == 0 {
		if len(req.Scope) > 0 {
			for receiverRange, receiverIds := range req.Scope {
				tmpSql := " OR (receiver_range = %d AND receiver_id IN (%s))"
				tmpSql = fmt.Sprintf(tmpSql, receiverRange, strings.Join(receiverIds, ","))
				scopeSql = scopeSql + tmpSql
			}
		}

		sql = " FROM %s WHERE receiver_range = 1" + scopeSql
		sql = fmt.Sprintf(sql, model.Message{}.TableName())

		countSql = "SELECT COUNT(*)" + sql
		err = db.Raw(countSql).Count(&count).Error
		if err != nil {
			logUtils.Errorf("count message error", zap.String("error:", err.Error()))
			return
		}

		selectSql = "SELECT *" + sql
		err = db.Raw(selectSql).Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
			Find(&messages).Error
		if err != nil {
			logUtils.Errorf("query message error", zap.String("error:", err.Error()))
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

		for k, message := range messages {
			if _, ok := messageReadMap[message.ID]; ok {
				messages[k].ReadStatus = 2
			} else {
				messages[k].ReadStatus = 1
			}
		}
	} else {
		if len(req.Scope) > 0 {
			for receiverRange, receiverIds := range req.Scope {
				tmpSql := " OR (m.receiver_range = %d AND m.receiver_id IN (%s))"
				tmpSql = fmt.Sprintf(tmpSql, receiverRange, strings.Join(receiverIds, ","))
				scopeSql = scopeSql + tmpSql
			}
		}

		sql = " FROM %s m LEFT JOIN %s r ON m.id=r.message_id WHERE (m.receiver_range = 1 %s ) AND r.id IS"
		//未读
		if req.ReadStatus == 1 {
			sql = sql + " NULL"
		} else if req.ReadStatus == 2 {
			//已读
			sql = sql + " NOT NULL"
		}
		sql = fmt.Sprintf(sql, model.Message{}.TableName(), model.MessageRead{}.TableName(), scopeSql)

		countSql = "SELECT COUNT(*)" + sql
		err = db.Raw(countSql).Count(&count).Error
		if err != nil {
			logUtils.Errorf("count message error", zap.String("error:", err.Error()))
			return
		}

		selectSql = "SELECT m.*" + sql
		err = db.Raw(selectSql).Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
			Find(&messages).Error
		if err != nil {
			logUtils.Errorf("query message error", zap.String("error:", err.Error()))
			return
		}

		for k, _ := range messages {
			messages[k].ReadStatus = req.ReadStatus
		}
	}

	data.Populate(messages, count, req.Page, req.PageSize)

	return
}

func (r *MessageRepo) GetUnreadCount(scope v1.MessageScope) (count int64, err error) {
	var scopeSql string

	if len(scope.Scope) > 0 {
		for receiverRange, receiverIds := range scope.Scope {
			tmpSql := " OR (m.receiver_range = %d AND m.receiver_id IN (%s))"
			tmpSql = fmt.Sprintf(tmpSql, receiverRange, strings.Join(receiverIds, ","))
			scopeSql = scopeSql + tmpSql
		}
	}

	sql := "SELECT COUNT(*) FROM %s m LEFT JOIN %s r ON m.id=r.message_id WHERE (m.receiver_range = 1 %s ) AND r.id IS NULL"
	sql = fmt.Sprintf(sql, model.Message{}.TableName(), model.MessageRead{}.TableName(), scopeSql)

	err = r.DB.Raw(sql).Count(&count).Error
	if err != nil {
		logUtils.Errorf("count unread message error", zap.String("error:", err.Error()))
		return
	}
	return
}
