package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
)

type ProjectRoleMenuRepo struct {
	*BaseRepo       `inject:""`
	DB              *gorm.DB         `inject:""`
	ProjectRoleRepo *ProjectRoleRepo `inject:""`
	ProjectMenuRepo *ProjectMenuRepo `inject:""`
}

func (r *ProjectRoleMenuRepo) FindByRoleAndMenu(tenantId consts.TenantId, roleId, menuId uint) (projectRoleMenu model.ProjectRoleMenu, err error) {
	db := r.GetDB(tenantId).Model(&model.ProjectRoleMenu{}).Where("role_id = ?", roleId).Where("menu_id = ?", menuId)

	err = db.First(&projectRoleMenu).Error
	return
}

func (r *ProjectRoleMenuRepo) Create(tenantId consts.TenantId, projectRoleMenu model.ProjectRoleMenu) (err error) {
	roleMenu, err := r.FindByRoleAndMenu(tenantId, projectRoleMenu.RoleId, projectRoleMenu.MenuId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logUtils.Errorf("创建角色菜单失败%s", err.Error())
		return
	}

	if roleMenu.ID != 0 {
		logUtils.Infof("角色id：%d的菜单id:%d已经存在", projectRoleMenu.RoleId, projectRoleMenu.MenuId)
		return
	}

	err = r.GetDB(tenantId).Create(&projectRoleMenu).Error
	if err != nil {
		logUtils.Errorf("创建角色菜单失败%s", err.Error())
		return
	}

	return
}

func (r *ProjectRoleMenuRepo) DeleteById(tenantId consts.TenantId, id uint) error {
	err := r.GetDB(tenantId).Unscoped().Delete(&model.ProjectRoleMenu{}, id).Error
	if err != nil {
		logUtils.Errorf("delete project role menu by id get  err ", zap.String("错误:", err.Error()))
		return err
	}
	return nil
}

func (r *ProjectRoleMenuRepo) GetRoleMenuConfig() (roleMenuConfigs []v1.ProjectRoleMenuConfig, err error) {

	data, err := ioutil.ReadFile("config/sample/role-menu.json")
	if err != nil {
		logUtils.Errorf("load role menu config err ", zap.String("错误:", err.Error()))
		return
	}
	roleMenuConfigs = make([]v1.ProjectRoleMenuConfig, 0)
	err = json.Unmarshal(data, &roleMenuConfigs)
	if err != nil {
		logUtils.Errorf("unmarshall role menu config err ", zap.String("错误:", err.Error()))
		return
	}
	return
}

func (r *ProjectRoleMenuRepo) GetConfigData(tenantId consts.TenantId) (menus []model.ProjectRoleMenu, err error) {
	roleMenuConfigs, err := r.GetRoleMenuConfig()
	if err != nil {
		logUtils.Errorf("load role menu config err ", zap.String("错误:", err.Error()))
		return
	}

	roleNameIdMap, err := r.ProjectRoleRepo.GetAllRoleNameIdMap(tenantId)
	if err != nil {
		logUtils.Errorf("get all role name id map err ", zap.String("错误:", err.Error()))
		return
	}

	menuCodeIdMap, err := r.ProjectMenuRepo.GetAllMenuCodeIdMap()
	if err != nil {
		logUtils.Errorf("get all menu code id map err ", zap.String("错误:", err.Error()))
		return
	}

	for _, config := range roleMenuConfigs {
		var roleId uint
		if _, ok := roleNameIdMap[config.RoleName]; ok {
			roleId = roleNameIdMap[config.RoleName]
		}
		for _, menuCode := range config.Menus {
			projectRoleMenu := model.ProjectRoleMenu{RoleId: roleId}
			if menuId, ok := menuCodeIdMap[menuCode]; ok {
				projectRoleMenu.MenuId = menuId
				menus = append(menus, projectRoleMenu)
			} else {
				continue
			}
		}
	}

	return
}

func (r *ProjectRoleMenuRepo) DeleteAllData(tenantId consts.TenantId) {
	r.GetDB(tenantId).Delete(&model.ProjectRoleMenu{}, "id > 0")
}

func (r *ProjectRoleMenuRepo) BatchCreate(tenantId consts.TenantId, roleMenus []model.ProjectRoleMenu) (successCount int, failItems []string) {
	var err error
	for _, roleMenu := range roleMenus {
		err = r.Create(tenantId, roleMenu)
		if err != nil {
			failItems = append(failItems, fmt.Sprintf("为角色%d添加菜单%d失败，错误%s", roleMenu.RoleId, roleMenu.MenuId, err.Error()))
		} else {
			successCount++
		}
	}
	return
}
