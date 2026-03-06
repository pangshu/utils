package converter

import (
	"math/big"
)

func (c *Converter[T]) ToBigFloatOrDefault(v T, defaultValue *big.Float) *big.Float {
	result, err := c.ToBigFloat(v)
	if err != nil {
		return defaultValue
	}
	return result
}
