package jslibHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/dop251/goja_nodejs/require"
	"path/filepath"
	"sync"
	"time"
)

var (
	ServerLoadedLibs sync.Map
)

func LoadServerJslibs(require *require.RequireModule) {
	LoadCacheIfNeeded()

	JslibCache.Range(func(key, value interface{}) bool {
		id := key.(uint)

		lib := value.(Jslib)

		updateTime, ok := GetServerCache(id)
		if !ok || updateTime.Before(lib.UpdatedAt) {
			pth := filepath.Join(consts.WorkDir, fmt.Sprintf("%d.js", id))
			fileUtils.WriteFile(pth, lib.Script)
			require.Require(pth)
		}

		return true
	})
}

func GetServerCache(id uint) (val time.Time, ok bool) {
	inf, ok := ServerLoadedLibs.Load(id)

	if ok {
		val = inf.(time.Time)
	}

	return
}

func SetServerCache(id uint, val time.Time) {
	ServerLoadedLibs.Store(id, val)
}
