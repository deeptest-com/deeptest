package jslibHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func LoadServerJslibs(tenantId consts.TenantId, runtime *goja.Runtime, require *require.RequireModule, projectId uint) {
	LoadCacheIfNeeded(tenantId)

	JslibCache.Range(func(key, value interface{}) bool {

		lib, ok := value.(Jslib)
		if !ok {
			return true
		}

		if lib.ProjectId != projectId {
			return true
		}

		tmpFile := fmt.Sprintf("%s-%d.js", key, lib.UpdatedAt.Unix())
		tmpPath := fmt.Sprintf("%s/%s.js", consts.TmpDirRelativeServer, tmpFile)
		tmpContent := lib.Script
		fileUtils.WriteFileIfNotExist(tmpPath, tmpContent)

		module, err := require.Require("./" + tmpPath)
		if err != nil {
			logUtils.Infof("goja require failed, path: %s, err: %s.", tmpPath, err.Error())
			return true
		}

		runtime.Set(lib.Name, module)

		res := module.Export()
		functions, ok := res.(map[string]interface{})
		if ok {
			lib.Functions = make([]JsFunc, 0)
			for funcName, _ := range functions {
				var args string
				args, err = GetJsFunsParams(runtime, fmt.Sprintf("%s.%s", lib.Name, funcName))
				if err == nil {
					lib.Functions = append(lib.Functions, JsFunc{Name: funcName, Args: args})
				}

			}
			SetJslibCache(tenantId, lib.Id, lib)
		}

		return true
	})
}

func GetJsFunsParams(runtime *goja.Runtime, funName string) (res string, err error) {
	script := `function getFunctionParameters(func) {  
    const str = func.toString();  
    const start = str.indexOf('(') + 1;  
    const end = str.indexOf(')');  
  
    if (start < 0 || end < 0) {  
        return []; // 如果函数没有参数或格式不正确，返回空数组  
    }  
  
    const result = str.slice(start, end).match(/\b\w+\b/g); // 使用正则表达式匹配参数名称  
    if (result === null) {  
        return []; // 如果没有匹配到参数，返回空数组  
    }  
  
    return result;  
}  
getFunctionParameters(` + funName + `)`
	value, err := runtime.RunString(script)
	return value.String(), err

}
