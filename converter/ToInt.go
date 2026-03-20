package converter

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt converts numeric types to int.
func (c *Converter[T]) ToInt(v T) (int, error) {
	if true == c.IsNil(v) {
		return 0, nil
	}

	switch val := any(v).(type) {
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		if val >= math.MinInt && val <= math.MaxInt {
			return 0, nil
		} else {
			return int(val), fmt.Errorf("int64 value %d overflows int", v)
		}
	case uint:
		if val <= math.MaxInt {
			return int(val), nil
		} else {
			return int(val), fmt.Errorf("uint value %d overflows int", v)
		}
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		intSize := 32 << (^uint(0) >> 63) // 32 or 64
		if math.MaxUint32 >= (1<<(intSize-1) - 1) {
			if val <= (1<<(intSize-1) - 1) {
				return 0, nil
			} else {
				return int(val), fmt.Errorf("uint32 value %d overflows int", v)
			}
		} else {
			return int(val), nil
		}
	case uint64:
		if val <= math.MaxInt {
			return int(val), nil
		} else {
			return int(val), fmt.Errorf("uint32 value %d overflows int", v)
		}
	case float32:
		if val <= math.MaxInt {
			return int(val), nil
		} else {
			return int(val), fmt.Errorf("float32 value %d overflows int", v)
		}
	case float64:
		if val <= math.MaxInt {
			return int(val), nil
		} else {
			return int(val), fmt.Errorf("float64 value %d overflows int", v)
		}
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		tmpVal, err := strconv.ParseInt(string(val), 0, 0)
		if err != nil {
			return 0, fmt.Errorf("[]byte value %d overflows int", v)
		}
		return int(tmpVal), nil
	case string:
		tmpVal, err := strconv.ParseInt(val, 0, 0)
		if err != nil {
			return 0, fmt.Errorf("string value %d overflows int", v)
		}
		return int(tmpVal), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T for ToInt", val)
	}
}
