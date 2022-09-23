package cacheUtils

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/kataras/iris/v12"
)

func GetAllExtractedVariablesForProject(projectId string) (ret []serverDomain.Variable) {
	mapObj, ok := consts.ExtractedVariableCache.Load(projectId)
	if !ok {
		return
	}

	for k, v := range mapObj.(iris.Map) {
		variable := serverDomain.Variable{
			Name:  k,
			Value: fmt.Sprintf("%v", v),
		}
		ret = append(ret, variable)
	}

	return
}

func GetExtractedVariableFromCache(projectId, key string) (val string) {
	mapObj, ok := consts.ExtractedVariableCache.Load(projectId)

	var mp iris.Map
	if ok {
		mp = mapObj.(iris.Map)
	} else {
		mp = iris.Map{}
	}

	valObj, _ := mp[key]
	if valObj != nil {
		val = valObj.(string)
	}

	return
}

func SetExtractedVariableToCache(projectId, key, val string) {
	mapObj, ok := consts.ExtractedVariableCache.Load(projectId)

	var mp iris.Map
	if ok {
		mp = mapObj.(iris.Map)
	} else {
		mp = iris.Map{}
	}

	mp[key] = val

	consts.ExtractedVariableCache.Store(projectId, mp)

	return
}

func ClearExtractedVariables(projectId string) {
	mp := iris.Map{}
	consts.ExtractedVariableCache.Store(projectId, mp)
}
