package casesHelper

import "math/rand"

const (
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandStr() (ret string) {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	ret = string(b)

	return
}

func RandInt32() (ret int32) {
	ret = rand.Int31()
	return
}

func RandInt64() (ret int64) {
	ret = rand.Int63()
	return
}

func RandFloat32() (ret float32) {
	ret = rand.Float32()
	return
}

func RandFloat64() (ret float64) {
	ret = rand.Float64()
	return
}

func RandBool() (ret bool) {
	r := rand.Intn(1)
	if r == 0 {
		ret = true
	} else {
		ret = false
	}

	return
}
