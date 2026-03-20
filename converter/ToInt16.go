package converter

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt16 converts numeric types to int16.
func (c *Converter[T]) ToInt16(v T) (int16, error) {
	if true == c.IsNil(v) {
		return 0, nil
	}

	var tmpVal int64
	var err error

	switch val := any(v).(type) {
	case int:
		tmpVal = int64(val)
	case int8:
		tmpVal = int64(val)
	case int16:
		tmpVal = int64(val)
	case int32:
		tmpVal = int64(val)
	case int64:
		tmpVal = val
	case uint:
		tmpVal = int64(val)
	case uint8:
		tmpVal = int64(val)
	case uint16:
		tmpVal = int64(val)
	case uint32:
		tmpVal = int64(val)
	case uint64:
		if val > uint64(math.MaxInt64) {
			return 0, fmt.Errorf("uint64 value %d overflows int64", val)
		}
		tmpVal = int64(val)
	case float32:
		tmpVal = int64(val)
	case float64:
		tmpVal = int64(val)
	case bool:
		if val {
			tmpVal = 1
		} else {
			tmpVal = 0
		}
	case []byte:
		tmpVal, err = strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("[]byte value %d overflows int64", v)
		}
	case string:
		tmpVal, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("[]byte value %d overflows int64", val)
		}
	default:
		return 0, fmt.Errorf("unsupported type: %T for ToInt16", val)
	}

	if tmpVal < math.MinInt16 || tmpVal > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflows int64", tmpVal)
	}
	return int16(tmpVal), nil
}
