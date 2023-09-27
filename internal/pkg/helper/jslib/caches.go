package jslibHelper

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/snowlyg/helper/dir"
	"path/filepath"
	"sync"
)

var (
	IsCacheReady bool
	JslibCache   sync.Map
)

func LoadCacheIfNeeded() (err error) {
	if IsCacheReady {
		return
	}

	InitJslibCache()
	IsCacheReady = true

	return
}

func GetJslibCache(id uint) (val Jslib) {
	inf, ok := JslibCache.Load(id)

	if ok {
		val = inf.(Jslib)
	}

	return
}

func SetJslibCache(id uint, val Jslib) {
	JslibCache.Store(id, val)
}

func InitJslibCache() (err error) {
	db := dao.GetDB()
	if db == nil {
		return
	}

	var pos []SysJslib
	err = db.Model(&SysJslib{}).
		Where("NOT deleted").
		Find(&pos).Error

	for _, po := range pos {
		pth := filepath.Join(dir.GetCurrentAbPath(), po.ScriptFile)
		content := fileUtils.ReadFile(pth)

		to := Jslib{
			Script:    content,
			UpdatedAt: *po.UpdatedAt,
		}
		if po.UpdatedAt != nil {
			to.UpdatedAt = *po.UpdatedAt
		} else {
			to.UpdatedAt = *po.CreatedAt
		}

		SetJslibCache(po.ID, to)
	}

	return
}
