package cacheUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

func GetCache(scope, key string) (val string) {
	mapObj, ok := consts.CacheData.Load(scope)

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
	mapObj, ok := consts.CacheData.Load(scope)

	var mp iris.Map
	if ok {
		mp = mapObj.(iris.Map)
	} else {
		mp = iris.Map{}
	}

	mp[key] = val

	consts.CacheData.Store(scope, mp)

	return
}
