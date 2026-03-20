package converter

import (
	"math/big"
)

// ToBigRatOrDefault 将数值类型转换为*big.Rat，失败就返回默认值。 / converts numeric types to *big.Rat, returns defaultValue on failure.
func (c *Converter[T]) ToBigRatOrDefault(v T, defaultValue *big.Rat) *big.Rat {
	result, err := c.ToBigRat(v)
	if err != nil {
		return defaultValue
	}
	return result
}
