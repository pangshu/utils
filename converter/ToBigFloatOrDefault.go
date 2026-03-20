package converter

import (
	"math/big"
)

// ToBigFloatOrDefault 将数值类型转换为*big.Float，失败就返回默认值。 / converts numeric types  to *big.Float, returns defaultValue on failure.
func (c *Converter[T]) ToBigFloatOrDefault(v T, defaultValue *big.Float) *big.Float {
	result, err := c.ToBigFloat(v)
	if err != nil {
		return defaultValue
	}
	return result
}
