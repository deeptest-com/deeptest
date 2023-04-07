package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"gorm.io/gorm"
	"strings"
)

type ProjectRoleMenuRepo struct {
	DB *gorm.DB `inject:""`
}

func NewProjectRoleMenuRepo() *ProjectRoleMenuRepo {
	return &ProjectRoleMenuRepo{}
}

func (r *ProjectRoleMenuRepo) GetRoleMenuList(roleId uint) (roleMenus []model.ProjectRoleMenu, err error) {
	err = r.DB.Model(&model.ProjectRoleMenu{}).
		Where("project_role_id = ?", roleId).
		Scan(&roleMenus).Error

	return
}

func (r *ProjectRoleMenuRepo) GetMenuDescByName(name string) string {
	desc := ""
	menuMap := map[string]string{
		"projects":     "项目",
		"datapools":    "数据池",
		"interfaces":   "接口",
		"environments": "环境",
		"extractors":   "提取器",
		"scenarios":    "场景",
		"plans":        "计划",
	}
	if value, ok := menuMap[name]; ok {
		desc = value
	}
	return desc
}

func (r *ProjectRoleMenuRepo) GetMenusForRole() (res []model.ProjectRoleMenu, err error) {
	var projectRolePerms []model.ProjectRolePerm
	err = r.DB.Model(&model.ProjectRolePerm{}).Find(&projectRolePerms).Error
	if err != nil {
		return
	}

	permsMap := make(map[uint]uint)
	for _, v := range projectRolePerms {
		permsMap[v.ProjectPermId] = v.ProjectPermId
	}

	var permsArr []uint
	for _, v := range permsMap {
		permsArr = append(permsArr, v)
	}

	var projectPerms []model.ProjectPerm
	err = r.DB.Model(&model.ProjectPerm{}).Where("id IN (?)", permsArr).Find(&projectPerms).Error
	if err != nil {
		return
	}

	permMenuMap := make(map[uint]model.ProjectRoleMenu)
	for _, v := range projectPerms {
		if !strings.Contains(v.Name, "/api/v1/") {
			continue
		}
		pathArr := strings.Split(v.Name, "/")
		menuName := pathArr[3]
		menuDesc := r.GetMenuDescByName(menuName)
		projectRoleMenuTmp := model.ProjectRoleMenu{
			ProjectMenuName:        menuName,
			ProjectMenuDescription: menuDesc,
		}
		permMenuMap[v.ID] = projectRoleMenuTmp
	}

	roleMenusMap := make(map[uint][]string)
	for _, v := range projectRolePerms {
		if roleMenu, ok := permMenuMap[v.ProjectPermId]; ok {
			roleMenusMap[v.ProjectRoleId] = append(roleMenusMap[v.ProjectRoleId], roleMenu.ProjectMenuName)
		}
	}

	for k, v := range roleMenusMap {
		roleMenusMap[k] = _commUtils.ArrayRemoveDuplication(v)
	}

	for k, v := range roleMenusMap {
		for _, v1 := range v {
			menuDesc := r.GetMenuDescByName(v1)
			projectRoleMenuTmp := model.ProjectRoleMenu{
				ProjectRoleId:          k,
				ProjectMenuName:        v1,
				ProjectMenuDescription: menuDesc,
			}
			res = append(res, projectRoleMenuTmp)
		}
	}

	return
}

func (r *ProjectRoleMenuRepo) BatchAddData(menus []model.ProjectRoleMenu) (successCount int, failItems []string) {
	for _, menu := range menus {
		err := r.DB.Model(&model.ProjectRoleMenu{}).Create(&menu).Error
		if err != nil {
			failItems = append(failItems, fmt.Sprintf("为角色%d添加菜单%s失败，错误%s", menu.ProjectRoleId, menu.ProjectMenuName, err.Error()))
		} else {
			successCount++
		}
	}
	return
}
