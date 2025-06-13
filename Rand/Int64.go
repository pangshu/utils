package Rand

import (
	"math/rand"
	"time"
)

// Int64 随机区间值
func (*Rand) Int64(min, max int64) int64 {
	// 防止出错，对min与max进行互换
	if min > max {
		min, max = max, min
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Int63n(max-min+1)
}
