package util

import (
	"math/rand"
	"time"
)

func GetRandomDuration(minSec, maxSec int) time.Duration {
	if minSec >= maxSec {
		return -1
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := (r.Intn(maxSec*1000) - minSec*1000) * int(time.Millisecond)

	return time.Duration(ret)
}
