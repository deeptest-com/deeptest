package jslibHelper

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	fileUtils "github.com/deeptest-com/deeptest/pkg/lib/file"
	"path/filepath"
	"sync"
)

var (
	IsCacheReady bool
	JslibCache   sync.Map
)

func LoadCacheIfNeeded(tenantId consts.TenantId) (err error) {
	if IsCacheReady {
		return
	}

	InitJslibCache(tenantId)
	IsCacheReady = true

	return
}

func GetJslibCache(tenantId consts.TenantId, id uint) (val Jslib) {
	key := fmt.Sprintf("%s_%d", tenantId, id)
	inf, ok := JslibCache.Load(key)

	if ok {
		val = inf.(Jslib)
	}

	return
}

func SetJslibCache(tenantId consts.TenantId, id uint, val Jslib) {
	key := fmt.Sprintf("%s_%d", tenantId, id)
	JslibCache.Store(key, val)
}

func InitJslibCache(tenantId consts.TenantId) (err error) {
	db := dao.GetDB(tenantId)
	if db == nil {
		return
	}

	var pos []SysJslib
	err = db.Model(&SysJslib{}).
		Where("NOT deleted").
		Find(&pos).Error

	for _, po := range pos {
		pth := filepath.Join(consts.WorkDir, po.ScriptFile)
		content := fileUtils.ReadFile(pth)

		to := Jslib{
			Id:        po.ID,
			ProjectId: po.ProjectId,
			Name:      po.Name,
			Script:    content,
			UpdatedAt: *po.UpdatedAt,
		}
		if po.UpdatedAt != nil {
			to.UpdatedAt = *po.UpdatedAt
		} else {
			to.UpdatedAt = *po.CreatedAt
		}

		SetJslibCache(tenantId, po.ID, to)
	}

	return
}
