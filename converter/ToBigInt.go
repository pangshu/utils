package converter

import (
	"fmt"
	"math"
	"math/big"
)

func (c *Converter[T]) ToBigInt(v T) (*big.Int, error) {
	result := new(big.Int)
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
		if val <= math.MaxInt {
			result.SetInt64(int64(val))
		} else {
			return nil, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case float64:
		if val <= math.MaxInt {
			result.SetInt64(int64(val))
		} else {
			return nil, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		_, ok := result.SetString(val, 10)
		if !ok {
			return nil, fmt.Errorf("failed to convert string to big.Int: %s", val)
		}
	default:
		return nil, fmt.Errorf("unsupported type: %T", val)
	}
	return result, nil
}
