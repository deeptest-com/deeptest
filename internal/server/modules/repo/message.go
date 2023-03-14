package repo

import (
	"gorm.io/gorm"
	"strconv"
)

type MessageRepo struct {
	DB          *gorm.DB     `inject:""`
	BaseRepo    *BaseRepo    `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func NewMessageRepo() *MessageRepo {
	return &MessageRepo{}
}

func (r *MessageRepo) GetScope(userId uint) (scope map[int][]string) {
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
