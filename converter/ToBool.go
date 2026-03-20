package converter

import (
	"fmt"
	"strconv"
)

// ToBool converts numeric types to bool.
func (c *Converter[T]) ToBool(v T) (bool, error) {
	if true == c.IsNil(v) {
		return false, fmt.Errorf("cannot convert nil to bool")
	}

	switch val := any(v).(type) {
	case int:
		return val != 0, nil
	case int8:
		return val != 0, nil
	case int16:
		return val != 0, nil
	case int32:
		return val != 0, nil
	case int64:
		return val != 0, nil
	case uint:
		return val != 0, nil
	case uint8:
		return val != 0, nil
	case uint16:
		return val != 0, nil
	case uint32:
		return val != 0, nil
	case uint64:
		return val != 0, nil
	case float32:
		return val != 0, nil
	case float64:
		return val != 0, nil
	case bool:
		return val, nil
	case []byte:
		return strconv.ParseBool(string(val))
	case string:
		return strconv.ParseBool(val)
	default:
		return false, fmt.Errorf("unsupported type: %T for ToBool", val)
	}
}
