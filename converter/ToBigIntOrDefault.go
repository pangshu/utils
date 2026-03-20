package converter

import (
	"math/big"
)

// ToBigIntOrDefault 将数值类型转换为*big.Int，失败就返回默认值。 / converts numeric types to *big.Int, returns defaultValue on failure.
func (c *Converter[T]) ToBigIntOrDefault(v T, defaultValue *big.Int) *big.Int {
	result, err := c.ToBigInt(v)
	if err != nil {
		return defaultValue
	}
	return result
}
