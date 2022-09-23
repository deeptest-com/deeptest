package extractCache

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"sync"
)

var (
	extractedVariableCache sync.Map
)

func Get(key string) (ret string) {
	obj, ok := extractedVariableCache.Load(key)
	if ok {
		ret = obj.(string)
	}

	return
}

func Set(key, val string) {
	extractedVariableCache.Store(key, val)
	return
}

func GetAll() (ret []serverDomain.Variable) {
	extractedVariableCache.Range(func(key, value interface{}) bool {
		variable := serverDomain.Variable{
			Name:  key.(string),
			Value: value.(string),
		}
		ret = append(ret, variable)

		return true
	})

	return
}

func ClearAll() {
	extractedVariableCache.Range(func(key, value interface{}) bool {
		extractedVariableCache.Delete(key)
		return true
	})
}
