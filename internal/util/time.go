package util

import (
	"math/rand"
	"time"
)

func GetRandomDuration(minSec, maxSec int, seedOffset int64) time.Duration {
	if minSec >= maxSec {
		return -1
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano() + seedOffset))
	ret := (r.Intn((maxSec-minSec)*1000) + minSec*1000) * int(time.Millisecond)

	return time.Duration(ret)
}
