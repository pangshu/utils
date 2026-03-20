package converter

import (
	"fmt"
	"strconv"
)

// ToFloat64 converts numeric types to Float64.
func (c *Converter[T]) ToFloat64(v T) (float64, error) {
	if true == c.IsNil(v) {
		return 0, nil
	}

	switch val := any(v).(type) {
	case int:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return strconv.ParseFloat(string(val), 64)
	case string:
		return strconv.ParseFloat(val, 64)
	default:
		return 0, fmt.Errorf("unsupported type: %T for ToFloat64", val)
	}
}
