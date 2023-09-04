package valueUtils

import (
	"math/rand"
	"time"
)

func RandNum(length int) int {
	randSeeds := time.Now().Unix() + int64(rand.Intn(100000000))
	rand.Seed(randSeeds)

	seedInt := rand.Intn(length)
	return seedInt
}
func RandNum64(length int64) int64 {
	randSeeds := time.Now().Unix() + int64(rand.Intn(100000000))
	rand.Seed(randSeeds)

	seedInt := rand.Int63n(length)
	return seedInt
}
