package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
)

type ProjectMenuRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ProjectMenuRepo) FindByCode(tenantId consts.TenantId, code string) (projectMenu model.ProjectMenu, err error) {
	db := r.GetDB(tenantId).Model(&model.ProjectMenu{}).Where("code = ?", code)

	err = db.First(&projectMenu).Error
	return
}

func (r *ProjectMenuRepo) Create(tenantId consts.TenantId, projectMenu model.ProjectMenu) (id uint, err error) {
	menu, err := r.FindByCode(tenantId, projectMenu.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logUtils.Errorf("创建项目菜单失败%s", err.Error())
		return
	}

	if menu.ID != 0 {
		logUtils.Infof("项目菜单%s已经存在", projectMenu.Code)
		return
	}

	err = r.GetDB(tenantId).Create(&projectMenu).Error
	if err != nil {
		logUtils.Errorf("创建项目菜单失败%s", err.Error())
		return
	}
	id = projectMenu.ID
	return
}

func (r *ProjectMenuRepo) Update(tenantId consts.TenantId, req v1.ProjectMenuReq) error {
	projectMenu := model.ProjectMenu{ProjectMenuBase: req.ProjectMenuBase}
	err := r.GetDB(tenantId).Model(&model.ProjectMenu{}).Where("id = ?", req.Id).Updates(&projectMenu).Error
	if err != nil {
		logUtils.Errorf("update project menu error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectMenuRepo) DeleteById(tenantId consts.TenantId, id uint) error {
	err := r.GetDB(tenantId).Unscoped().Delete(&model.ProjectMenu{}, id).Error
	if err != nil {
		logUtils.Errorf("delete project menu by id get  err ", zap.String("错误:", err.Error()))
		return err
	}
	return nil
}

func (r *ProjectMenuRepo) GetRoleMenuList(tenantId consts.TenantId, roleId uint) (roleMenus []model.ProjectMenu, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMenu{}).
		Joins("left join biz_project_role_menu m on biz_project_menu.id = m.menu_id").
		Where("m.role_id = ?", roleId).
		Scan(&roleMenus).Error

	return
}

func (r *ProjectMenuRepo) GetRoleMenuCodeList(tenantId consts.TenantId, roleId uint) (roleMenuCodes []string, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMenu{}).
		Joins("left join biz_project_role_menu m on biz_project_menu.id = m.menu_id").
		Select("biz_project_menu.code").
		Where("m.role_id = ?", roleId).
		Scan(&roleMenuCodes).Error

	return
}

func (r *ProjectMenuRepo) GetAllMenuList(tenantId consts.TenantId) (menus []model.ProjectMenu, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMenu{}).Scan(&menus).Error
	return
}

func (r *ProjectMenuRepo) GetMenuConfig(tenantId consts.TenantId) (menuConfigs []v1.ProjectMenuConfig, err error) {
	data, err := ioutil.ReadFile("config/sample/menu.json")
	if err != nil {
		logUtils.Errorf("load menu config err ", zap.String("错误:", err.Error()))
		return
	}
	menuConfigs = make([]v1.ProjectMenuConfig, 0)
	err = json.Unmarshal(data, &menuConfigs)
	fmt.Printf("menuConfigs:%+v, err:%+v \n", menuConfigs, err)

	if err != nil {
		logUtils.Errorf("unmarshall menu config err ", zap.String("错误:", err.Error()))
		return
	}
	return
}

func (r *ProjectMenuRepo) GetConfigData(tenantId consts.TenantId, level string) (menus []model.ProjectMenu, err error) {
	menuConfigs, err := r.GetMenuConfig(tenantId)
	if err != nil {
		return
	}

	for _, v := range menuConfigs {
		projectMenuBase := v1.ProjectMenuBase{
			Code:  v.Code,
			Title: v.Title,
			Path:  v.Path,
			Type:  v.Type,
		}

		if level == "firstLevel" {
			if v.Parent != "" {
				continue
			}
		} else {
			if level == "secondLevel" {
				if v.Parent == "" || v.Type == "button" {
					continue
				}
			} else {
				if v.Type != "button" {
					continue
				}
			}

			if v.Parent != "" {
				parentProjectMenu, err := r.FindByCode(tenantId, v.Parent)
				if err != nil {
					logUtils.Errorf("get parent menu err ", zap.String("错误:", err.Error()))
					continue
				}
				projectMenuBase.ParentId = parentProjectMenu.ID
			}
		}

		projectMenu := model.ProjectMenu{ProjectMenuBase: projectMenuBase}
		menus = append(menus, projectMenu)
	}
	return
}

func (r *ProjectMenuRepo) BatchCreate(tenantId consts.TenantId, menus []model.ProjectMenu) (successCount int, failItems []string) {
	for _, menu := range menus {
		_, err := r.Create(tenantId, menu)
		if err != nil {
			failItems = append(failItems, menu.Code)
			continue
		}
		successCount++
	}
	return
}

func (r *ProjectMenuRepo) DeleteAllData(tenantId consts.TenantId) {
	r.GetDB(tenantId).Delete(&model.ProjectMenu{}, "id > 0")
}

func (r *ProjectMenuRepo) BatchInitData(tenantId consts.TenantId, level string) (successCount int, failItems []string) {
	data, err := r.GetConfigData(tenantId, level)
	if err != nil {
		failItems = append(failItems, fmt.Sprintf("%s级别所有数据", level))
		return
	}
	return r.BatchCreate(tenantId, data)
}

func (r *ProjectMenuRepo) GetAllMenuCodeIdMap(tenantId consts.TenantId) (data map[string]uint, err error) {
	menuList, err := r.GetAllMenuList(tenantId)
	if err != nil {
		logUtils.Errorf("get all menu list err ", zap.String("错误:", err.Error()))
		return
	}
	data = make(map[string]uint)
	for _, v := range menuList {
		data[v.Code] = v.ID
	}
	return
}
