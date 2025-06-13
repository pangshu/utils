package Var

import (
	"fmt"
	"strconv"
)

// ToFloat32 强制将变量转换为float32类型.
func (conv *Var) ToFloat32(in any) (float32, error) {
	i := indirect(in)
	switch s := i.(type) {
	case float64:
		return float32(s), nil
	case float32:
		return s, nil
	case int:
		return float32(s), nil
	case int64:
		return float32(s), nil
	case int32:
		return float32(s), nil
	case int16:
		return float32(s), nil
	case int8:
		return float32(s), nil
	case uint:
		return float32(s), nil
	case uint64:
		return float32(s), nil
	case uint32:
		return float32(s), nil
	case uint16:
		return float32(s), nil
	case uint8:
		return float32(s), nil
	case []byte:
		v, err := strconv.ParseFloat(string(s), 32)
		if err == nil {
			return float32(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return float32(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
		}
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	}
}
