package jslibHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func LoadServerJslibs(tenantId consts.TenantId, runtime *goja.Runtime, require *require.RequireModule) {
	LoadCacheIfNeeded(tenantId)

	JslibCache.Range(func(key, value interface{}) bool {
		id := key.(uint)
		if tenantId == "" {
			tenantId = "NA"
		}

		lib, ok := value.(Jslib)
		if !ok {
			return true
		}

		tmpFile := fmt.Sprintf("%d-%s-%d.js", id, tenantId, lib.UpdatedAt.Unix())
		tmpPath := fmt.Sprintf("%s/%s.js", consts.TmpDirRelativeServer, tmpFile)
		tmpContent := lib.Script
		fileUtils.WriteFileIfNotExist(tmpPath, tmpContent)

		module, err := require.Require("./" + tmpPath)
		if err != nil {
			logUtils.Infof("goja require failed, path: %s, err: %s.", tmpPath, err.Error())
		}

		runtime.Set(lib.Name, module)

		return true
	})
}
