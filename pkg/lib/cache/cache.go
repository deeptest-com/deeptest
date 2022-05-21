package _cacheUtils

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
)

func GetCache(scope, key string) (val string) {
	mapObj, ok := serverConsts.EnvVar.Load(scope)

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

func SetCache(scope, key, val string) {
	mapObj, ok := serverConsts.EnvVar.Load(scope)

	var mp iris.Map
	if ok {
		mp = mapObj.(iris.Map)
	} else {
		mp = iris.Map{}
	}

	mp[key] = val

	serverConsts.EnvVar.Store(scope, mp)

	return
}
