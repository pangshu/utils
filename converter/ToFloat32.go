package converter

import (
	"fmt"
	"math"
	"strconv"
)

// ToFloat32 converts numeric types to Float32.
func (c *Converter[T]) ToFloat32(v T) (float32, error) {
	if true == c.IsNil(v) {
		return 0, nil
	}

	switch val := any(v).(type) {
	case int:
		floatVal := float64(val)
		if math.Abs(floatVal) > math.MaxFloat32 {
			return 0, fmt.Errorf("int value %d overflows float32", val)
		}
		return float32(floatVal), nil
	case int8:
		return float32(val), nil
	case int16:
		return float32(val), nil
	case int32:
		return float32(val), nil
	case int64:
		floatVal := float64(val)
		if math.Abs(floatVal) > math.MaxFloat32 {
			return 0, fmt.Errorf("int64 value %d overflows float32", val)
		}
		return float32(floatVal), nil
	case uint:
		floatVal := float64(val)
		if math.Abs(floatVal) > math.MaxFloat32 {
			return 0, fmt.Errorf("int64 value %d overflows float32", val)
		}
		return float32(floatVal), nil
	case uint8:
		return float32(val), nil
	case uint16:
		return float32(val), nil
	case uint32:
		return float32(val), nil
	case uint64:
		floatVal := float64(val)
		if math.Abs(floatVal) > math.MaxFloat32 {
			return 0, fmt.Errorf("uint64 value %d overflows float32", val)
		}
		return float32(floatVal), nil
	case float32:
		return val, nil
	case float64:
		if math.Abs(val) > math.MaxFloat32 {
			return 0, fmt.Errorf("float64 value %f overflows float32", val)
		}
		return float32(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		floatVal, err := strconv.ParseFloat(string(val), 32)
		if err != nil {
			return 0, err
		}
		if math.Abs(floatVal) > math.MaxFloat32 {
			return 0, fmt.Errorf("[]byte value %f overflows float32", val)
		}
		return float32(floatVal), nil
	case string:
		floatVal, err := strconv.ParseFloat(string(val), 32)
		if err != nil {
			return 0, err
		}
		if math.Abs(floatVal) > math.MaxFloat32 {
			return 0, fmt.Errorf("[]byte value %f overflows float32", val)
		}
		return float32(floatVal), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T for ToFloat32", val)
	}
}
