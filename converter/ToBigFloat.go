package converter

import (
	"fmt"
	"math/big"
)

func (c *Converter[T]) ToBigFloat(v T) (*big.Float, error) {
	result := new(big.Float)
	switch val := any(v).(type) {
	case int:
		result.SetInt64(int64(val))
	case int8:
		result.SetInt64(int64(val))
	case int16:
		result.SetInt64(int64(val))
	case int32:
		result.SetInt64(int64(val))
	case int64:
		result.SetInt64(val)
	case uint:
		result.SetUint64(uint64(val))
	case uint8:
		result.SetUint64(uint64(val))
	case uint16:
		result.SetUint64(uint64(val))
	case uint32:
		result.SetUint64(uint64(val))
	case uint64:
		result.SetUint64(val)
	case float32:
		result.SetFloat64(float64(val))
	case float64:
		result.SetFloat64(val)
	case string:
		_, ok := result.SetString(val)
		if !ok {
			return nil, fmt.Errorf("failed to convert string to big.Float: %s", val)
		}
	default:
		return nil, fmt.Errorf("unsupported type: %T", val)
	}
	return result, nil
}
