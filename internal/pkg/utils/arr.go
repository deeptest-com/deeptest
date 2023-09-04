package commUtils

import (
	"math/rand"
	"time"
)

func RandArr(arr []interface{}) (ret []interface{}) {
	rand.Seed(time.Now().Unix())

	for range arr {
		rand := rand.Intn(len(arr))
		ret = append(ret, arr[rand])
	}

	return
}
