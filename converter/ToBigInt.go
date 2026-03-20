package converter

import (
	"fmt"
	"math"
	"math/big"
)

// ToBigInt converts numeric types to *big.Int.
func (c *Converter[T]) ToBigInt(v T) (*big.Int, error) {
	if true == c.IsNil(v) {
		return nil, fmt.Errorf("cannot convert nil to *big.Int")
	}

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
		if val <= math.MaxFloat32 {
			bigFloat := new(big.Float).SetFloat64(float64(val))
			bigInt, _ := bigFloat.Int(nil)
			result = bigInt
		} else {
			return nil, fmt.Errorf("float32 value is too large: %v", val)
		}
	case float64:
		if val <= math.MaxFloat64 {
			bigFloat := new(big.Float).SetFloat64(val)
			bigInt, _ := bigFloat.Int(nil)
			result = bigInt
		} else {
			return nil, fmt.Errorf("float64 value is too large: %v", val)
		}
	case bool:
		if val {
			result.SetInt64(1)
		} else {
			result.SetInt64(0)
		}
	case []byte:
		_, ok := result.SetString(string(val), 10)
		if !ok {
			return nil, fmt.Errorf("failed to convert []byte to *big.Int: %v", val)
		}
	case string:
		_, ok := result.SetString(val, 10)
		if !ok {
			return nil, fmt.Errorf("failed to convert string to *big.Int: %s", val)
		}
	default:
		return nil, fmt.Errorf("unsupported type: %T to ToBigInt", val)
	}
	return result, nil
}
