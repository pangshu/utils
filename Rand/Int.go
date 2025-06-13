package Rand

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
)

// Int 随机区间整数
func (*Rand) Int(min, max int) int {
	if min == max {
		return min
	}

	// 防止出错，对min与max进行互换
	if min > max {
		min, max = max, min
	}

	// 生成一个真随机的 uint64 数字
	var num uint64
	err := binary.Read(rand.Reader, binary.LittleEndian, &num)
	if err != nil {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
		return int(n.Int64() + int64(min))
	}

	// 将随机数映射到 [min, max] 范围内
	rangeSize := uint64(max - min + 1)
	maxRandom := math.MaxUint64 - (math.MaxUint64 % rangeSize)
	for num > maxRandom {
		err = binary.Read(rand.Reader, binary.LittleEndian, &num)
		if err != nil {
			n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
			return int(n.Int64() + int64(min))
		}
	}

	return int(num%rangeSize) + min
}
