package jslibHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"path/filepath"
	"sync"
	"time"
)

var (
	ServerLoadedLibs sync.Map
)

func LoadServerJslibs(runtime *goja.Runtime, require *require.RequireModule) {
	LoadCacheIfNeeded()

	JslibCache.Range(func(key, value interface{}) bool {
		id := key.(uint)

		lib, ok := value.(Jslib)
		if !ok {
			return true
		}

		updateTime, ok := GetServerCache(id)
		if !ok || updateTime.Before(lib.UpdatedAt) {
			pth := filepath.Join(consts.TmpDir, fmt.Sprintf("%d.js", id))
			fileUtils.WriteFile(pth, lib.Script)
			module, err := require.Require(pth)
			if err != nil {
				logUtils.Info(err.Error())
			}

			runtime.Set(lib.Name, module)

			SetServerCache(id, lib.UpdatedAt)
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
