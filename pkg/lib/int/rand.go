package _intUtils

import (
	"math/rand"
	"time"
)

func GenUniqueRandNum(start int, end int, count int) (ret []int) {
	if end < start || (end-start) < count {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(ret) < count {
		num := r.Intn(end-start) + start

		exist := false
		for _, v := range ret {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			ret = append(ret, num)
		}
	}

	return
}
