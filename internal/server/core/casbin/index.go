package casbin

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

var (
	once     sync.Once
	enforcer *casbin.Enforcer
	store    sync.Map
)

// Instance casbin 单例
func Instance(tenantId consts.TenantId) *casbin.Enforcer {
	res, _ := store.LoadOrStore(tenantId, GetEnforcer(tenantId))
	/*
		once.Do(func() {
			enforcer = GetEnforcer()
		})
	*/
	return res.(*casbin.Enforcer)
}

// GetEnforcer 获取 casbin.Enforcer
func GetEnforcer(tenantId consts.TenantId) *casbin.Enforcer {
	if dao.GetDB(tenantId) == nil {
		logUtils.Errorf("数据库未初始化")
		return nil
	}
	c, err := gormadapter.NewAdapterByDBUseTableName(dao.GetDB(tenantId), "", "sys_casbin_rule") // Your driver and data source.
	if err != nil {
		logUtils.Errorf("驱动初始化错误", zap.String("gormadapter.NewAdapterByDBUseTableName()", err.Error()))
		return nil
	}

	pth := filepath.Join(consts.WorkDir, consts.CasbinFileName) // created if needed in config init method
	enforcer, err := casbin.NewEnforcer(pth, c)
	if err != nil {
		logUtils.Errorf("初始化失败", zap.String("casbin.NewEnforcer()", err.Error()))
		return nil
	}

	if enforcer == nil {
		logUtils.Errorf("Casbin 未初始化")
		return nil
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logUtils.Errorf("加载规则失败", zap.String("casbin.LoadPolicy()", err.Error()))
		return nil
	}

	return enforcer
}

// GetRolesForUser 获取角色
func GetRolesForUser(tenantId consts.TenantId, uid uint) []string {
	uids, err := Instance(tenantId).GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		return []string{}
	}

	return uids
}

// GetPermissionsForUser 获取角色权限
func GetPermissionsForUser(tenantId consts.TenantId, id string) [][]string {
	return Instance(tenantId).GetPermissionsForUser(id)
}
