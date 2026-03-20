package converter

import (
	"fmt"
	"math/big"
)

// ToBigRat converts numeric types to *big.Rat.
func (c *Converter[T]) ToBigRat(v T) (*big.Rat, error) {
	if true == c.IsNil(v) {
		return nil, fmt.Errorf("cannot convert nil to *big.Rat")
	}

	result := new(big.Rat)
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
	case bool:
		if val {
			result.SetInt64(1)
		} else {
			result.SetInt64(0)
		}
	case []byte:
		_, ok := result.SetString(string(val))
		if !ok {
			return nil, fmt.Errorf("failed to convert []byte to *big.Rat: %v", val)
		}
	case string:
		_, ok := result.SetString(val)
		if !ok {
			return nil, fmt.Errorf("failed to convert string to *big.Rat: %s", val)
		}
	default:
		return nil, fmt.Errorf("unsupported type: %T to ToBigRat", val)
	}
	return result, nil
}
