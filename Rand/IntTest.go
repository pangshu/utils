package Rand

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
)

// TrueRandInt 生成一个真随机整数，范围在 [min, max] 之间
func TrueRandInt(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min should be less than or equal to max")
	}

	// 生成一个真随机的 uint64 数字
	var num uint64
	err := binary.Read(rand.Reader, binary.LittleEndian, &num)
	if err != nil {
		return 0, err
	}

	// 将随机数映射到 [min, max] 范围内
	rangeSize := uint64(max - min + 1)
	maxRandom := math.MaxUint64 - (math.MaxUint64 % rangeSize)
	for num > maxRandom {
		err = binary.Read(rand.Reader, binary.LittleEndian, &num)
		if err != nil {
			return 0, err
		}
	}

	return int(num%rangeSize) + min, nil
}
